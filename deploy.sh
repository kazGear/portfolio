#!/bin/bash

set -e

cd "$(dirname "$0")"

COMPOSE_CMD="docker compose \
  --env-file .env.prod \
  -f compose.base.yaml \
  -f compose.prod.yaml"

echo "=== Pull latest main ==="
git pull origin main

echo "=== Start DB ==="
$COMPOSE_CMD up -d db

echo "=== Generate sitemap ==="
$COMPOSE_CMD run --rm sitemap-generator sh -c "npm ci && npm run sitemap"

echo "=== Build and deploy ==="
$COMPOSE_CMD up --build --detach

echo "=== Container status ==="

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
echo "=== Deploy completed ==="