import tritonclient.http as httpclient
import logging

logger = logging.getLogger(__name__)

class TritonClient:
    def __init__(self, url, model_name, model_version, batch_size=1):
        self.url = url
        self.model_name = model_name
        self.model_version = model_version
        self.batch_size = batch_size
        self.client = self.initialize_client()
    
    def initialize_client(self):
        try:
            client = httpclient.InferenceServerClient(url=self.url, verbose=False)
            if not client.is_model_ready(model_name=self.model_name, model_version=self.model_version):
                raise ValueError(f"Model {self.model_name} version {self.model_version} is not ready.")
            logger.info("Successfully connected to Triton server and model is ready.")
            return client
        except Exception as e:
            logger.error(f"Failed to connect to Triton server: {e}")
            raise e
    
    def infer(self, inputs, outputs):
        try:
            response = self.client.infer(
                model_name=self.model_name,
                model_version=self.model_version,
                inputs=inputs,
                outputs=outputs
            )
            return response
        except Exception as e:
            logger.error(f"Inference error: {e}")
            raise e
