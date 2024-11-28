import random
import base64
import numpy as np
import logging
import tritonclient.http as httpclient
from flask import Flask, request, jsonify, render_template
from io import BytesIO
from PIL import Image

from config import Config
from utils.translator import PromptTranslator
from utils.image_handler import ImageHandler
from utils.client_ip import get_client_ip
from triton_client.client import TritonClient

app = Flask(__name__)

# Logging setup
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

# Load configuration file
app.config.from_object(Config)

# Initialize image handler
image_handler = ImageHandler(app.config['LOG_DIR'])

# Initialize translator
translator = PromptTranslator()

@app.route('/', methods=['GET'])
def index():
    return render_template('index.html')

@app.route('/infer/image', methods=['POST'])
def infer_image():
    try:
        data = request.get_json()
        if not data or 'prompt' not in data:
            logger.warning("Invalid request: 'prompt' is required.")

            return jsonify({"error": "Invalid request, 'prompt' is required."}), 400

        prompt = data['prompt']
        samples = 1
        steps = 45
        guidance_scale = 7.5
        seed = random.randint(0, 2**32 - 1)

        # Translation (Korean -> English)
        try:
            translated_prompt = translator.translate_prompt(prompt, src_lang='ko', dest_lang='en')
        except Exception as e:
            logger.error(f"Translation error: {e}")

            return jsonify({"error": f"Translation error: {e}"}), 500

        # Initialize Triton client
        try:
            triton_client = TritonClient(
                url=app.config['TRITON_URL'],
                model_name=app.config['MODEL_NAME'],
                model_version=app.config['MODEL_VERSION'],
                batch_size=app.config['BATCH_SIZE']
            )
        except Exception as e:
            logger.error(f"Triton client initialization error: {e}")

            return jsonify({"error": f"Triton client initialization error: {e}"}), 500

        inputs = []
        
        # PROMPT
        prompt_input = httpclient.InferInput("PROMPT", [app.config['BATCH_SIZE']], "BYTES")
        prompt_bytes = [translated_prompt.encode('utf-8')]
        prompt_input.set_data_from_numpy(np.array(prompt_bytes, dtype=object))
        inputs.append(prompt_input)
        
        # SAMPLES
        samples_input = httpclient.InferInput("SAMPLES", [app.config['BATCH_SIZE']], "INT32")
        samples_input.set_data_from_numpy(np.array([samples], dtype=np.int32))
        inputs.append(samples_input)
        
        # STEPS
        steps_input = httpclient.InferInput("STEPS", [app.config['BATCH_SIZE']], "INT32")
        steps_input.set_data_from_numpy(np.array([steps], dtype=np.int32))
        inputs.append(steps_input)
        
        # GUIDANCE_SCALE
        guidance_scale_input = httpclient.InferInput("GUIDANCE_SCALE", [app.config['BATCH_SIZE']], "FP32")
        guidance_scale_input.set_data_from_numpy(np.array([guidance_scale], dtype=np.float32))
        inputs.append(guidance_scale_input)
        
        # SEED
        seed_input = httpclient.InferInput("SEED", [app.config['BATCH_SIZE']], "INT64")
        seed_input.set_data_from_numpy(np.array([seed], dtype=np.int64))
        inputs.append(seed_input)

        outputs = [httpclient.InferRequestedOutput("IMAGES", binary_data=False)]

        # Inference Request
        try:
            response = triton_client.infer(inputs=inputs, outputs=outputs)
        except Exception as e:
            logger.error(f"Inference request error: {e}")
            return jsonify({"error": f"Inference request error: {e}"}), 500

        # Fetch image data
        images = response.as_numpy("IMAGES")
        if images is None:
            logger.error("No image returned from Triton server.")
            return jsonify({"error": "No image returned from Triton server."}), 500

        image_data = images[0]

        # Get user ID
        user_id = data['user']

        # Save image locally
        saved_image_path = image_handler.save_image_locally(user_id, image_data)
        if not saved_image_path:
            logger.error("Failed to save image locally.")
            return jsonify({"error": "Failed to save image locally."}), 500

        if image_data.dtype != np.uint8:
            image_data = (image_data * 255).astype(np.uint8)

        image = Image.fromarray(image_data)

        # Convert to PNG
        buffered = BytesIO()
        image.save(buffered, format="PNG")
        img_bytes = buffered.getvalue()

        # Base64 encoding
        img_base64 = base64.b64encode(img_bytes).decode('utf-8')

        logger.info("Image generation and saving successful.")

        return jsonify({"image": img_base64})

    except Exception as e:
        logger.error(f"Error in infer_image: {e}")

        return jsonify({"error": str(e)}), 500
