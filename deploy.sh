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
docker compose \
  --env-file .env.prod \
  -f compose.base.yaml \
  -f compose.prod.yaml \
  ps -a

echo "=== Deploy completed ==="
