#!/bin/sh

MONGO_CONTAINER_NAME="caju-code-challenge_db"
APP_CONTAINER_NAME="caju-code-challenge_api"

echo "Staring containers...\n"
docker-compose up --build -d

echo "-- Running unit tests --\n"
docker exec -it $APP_CONTAINER_NAME /usr/local/go/bin/go test ./tests/unit/... -v

echo "-- Running integration tests --\n"
docker exec -it $APP_CONTAINER_NAME /usr/local/go/bin/go test ./tests/integration/... -count=1 -v

echo "Cleaning...\n"
docker rm $MONGO_CONTAINER_NAME $APP_CONTAINER_NAME -f || true