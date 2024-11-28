#!/bin/bash

if [! -d "$PWD/models"]; then
    bash stable_diffusion.sh
fi

# Variables
SERVICE_NAME="triton-server"
SUBNET="103.0.0.0/24"
NETWORK="ai-artist"
IP="103.0.0.4"

# Start
docker network create --subnet $SUBNET $NETWORK

docker build -t $SERVICE_NAME . || { exit 1; }
docker run --rm $1 $2 $3 --name $SERVICE_NAME \
    --gpus all \
    --network $NETWORK \
    --ip $IP \
    -v $PWD/models:/models \
    $SERVICE_NAME \
    tritonserver --model-repository /models || { exit 1; }