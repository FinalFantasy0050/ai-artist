import os
from PIL import Image
import numpy as np
import logging

logger = logging.getLogger(__name__)

class ImageHandler:
    def __init__(self, log_dir):
        self.log_dir = log_dir
        self.ensure_log_dir()
    
    def ensure_log_dir(self):
        if not os.path.exists(self.log_dir):
            try:
                os.makedirs(self.log_dir)
                logger.info(f"Created log directory at {self.log_dir}")
            except Exception as e:
                logger.error(f"Failed to create log directory: {e}")
                
                raise e
    
    def save_image_locally(self, client_ip, image_data):
        try:
            client_dir = os.path.join(self.log_dir, client_ip)
            
            # Create the folder if it doesn't exist
            if not os.path.exists(client_dir):
                os.makedirs(client_dir)
                logger.info(f"Created directory for IP {client_ip} at {client_dir}")
    
            # Count the images to determine a unique number
            existing_files = [f for f in os.listdir(client_dir) if f.endswith('.png')]
            next_num = len(existing_files) + 1
            image_filename = f"{next_num}.png"
            image_path = os.path.join(client_dir, image_filename)
    
            # Convert the data format (float32 -> uint8)
            if image_data.dtype != np.uint8:
                image_data = (image_data * 255).astype(np.uint8)
    
            image = Image.fromarray(image_data)
    
            # Save the image
            image.save(image_path, format="PNG")
            logger.info(f"* Image saved to {image_path}")

            return image_path
        except Exception as e:
            logger.error(f"Failed to save image locally: {e}")

            return None
