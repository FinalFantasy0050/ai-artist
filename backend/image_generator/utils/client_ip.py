import logging
from flask import request

logger = logging.getLogger(__name__)

def get_client_ip():
    # Fetching the client's IP address.
    if request.headers.get('X-Forwarded-For'):
        ip = request.headers.get('X-Forwarded-For').split(',')[0].strip()
    else:
        ip = request.remote_addr
    logger.info(f"* Client IP: {ip}")

    return ip
