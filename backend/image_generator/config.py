import os

class Config:
    # Triton setting
    TRITON_URL = os.getenv("TRITON_URL", "103.0.0.4:8000")
    MODEL_NAME = os.getenv("MODEL_NAME", "stable_diffusion")
    MODEL_VERSION = os.getenv("MODEL_VERSION", "1")
    BATCH_SIZE = int(os.getenv("BATCH_SIZE", 1))
    
    # log folder path
    LOG_DIR = os.getenv("LOG_DIR", os.path.join(os.getcwd(), 'log'))
