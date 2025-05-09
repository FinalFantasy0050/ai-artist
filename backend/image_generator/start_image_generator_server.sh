#!/bin/bash

# Variables
SERVICE_NAME="image-generator-server"
SUBNET="103.0.0.0/24"
NETWORK="ai-artist"
IP="103.0.0.3"

# Start
docker network create --subnet $SUBNET $NETWORK

docker build -t $SERVICE_NAME . || { exit 1; }
docker run --rm $1 $2 $3 --name $SERVICE_NAME \
    --gpus all \
    --network $NETWORK \
    --ip $IP \
    -v $PWD/log:/app/log \
    $SERVICE_NAME || { exit 1; }