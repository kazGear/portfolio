#!/bin/bash

set -e

cd "$(dirname "$0")"

echo "=== Pull latest main ==="
git pull origin main

echo "=== Build and deploy ==="
docker compose \
  --env-file .env.prod \
  -f compose.base.yaml \
  -f compose.prod.yaml \
  up --build --detach

echo "=== Container status ==="

COMPOSE_CMD="docker compose \
  --env-file .env.prod \
  -f compose.base.yaml \
  -f compose.prod.yaml"

echo "=== Wait for containers to start ==="
sleep 10

$COMPOSE_CMD ps -a

echo "=== Check container status ==="

CONTAINERS=$($COMPOSE_CMD ps -aq)

ERROR=0

for CONTAINER in $CONTAINERS; do
    NAME=$(docker inspect -f '{{.Name}}' "$CONTAINER" | sed 's|^/||')
    STATUS=$(docker inspect -f '{{.State.Status}}' "$CONTAINER")

    echo "$NAME: $STATUS"

    if [ "$STATUS" != "running" ]; then
        echo "ERROR: $NAME is not running."
        ERROR=1
    fi
done

if [ "$ERROR" -ne 0 ]; then
    echo "=== Deploy failed: container error detected ==="
    exit 1
fi

echo "=== All containers are running ==="
echo "=== Deploy completed === "