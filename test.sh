#!/bin/sh

MONGO_CONTAINER_NAME="caju-code-challenge_db"
APP_CONTAINER_NAME="caju-code-challenge_api"

echo "Cleaning...\n"
docker rm $MONGO_CONTAINER_NAME $APP_CONTAINER_NAME -f || true

docker-compose up --build -d
sleep 10

echo "\n-- Running unit tests --"
docker exec -it $APP_CONTAINER_NAME /usr/local/go/bin/go test ./tests/unit/...

echo "-- Running integration tests --"
docker exec -it $APP_CONTAINER_NAME /usr/local/go/bin/go test ./tests/integration/...
