#!/bin/bash

# Variables
SERVICE_NAME="gateway-server"
SUBNET="103.0.0.0/24"
NETWORK="ai-artist"
IP="103.0.0.2"

# Start
docker network create --subnet $SUBNET $NETWORK

docker build -t $SERVICE_NAME . || { exit 1; }
docker run --rm $1 $2 $3 --name $SERVICE_NAME \
    --gpus all \
    --network $NETWORK \
    --ip $IP \
    -v $PWD/log:/app/log \
    -v $PWD/key.pem:/app/key.pem \
    -v $PWD/cert.pem:/app/cert.pem \
    -v $PWD/public:/app/public \
    -p 443:443 \
    $SERVICE_NAME || { exit 1; }