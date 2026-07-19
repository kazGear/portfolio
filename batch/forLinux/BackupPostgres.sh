#!/bin/bash

set -e

BACKUP_DIR="/home/kazuki/app/portfolio/infrastructure/db/backup"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="${BACKUP_DIR}/kaz_app_${DATE}.dump"

CONTAINER_NAME="portfolio-db-1"
DB_USER="postgres"
DB_NAME="kaz_app"

mkdir -p "$BACKUP_DIR"

echo "[$(date '+%Y-%m-%d %H:%M:%S')] PostgreSQL backup started."

docker exec "$CONTAINER_NAME" \
    pg_dump \
    -U "$DB_USER" \
    -d "$DB_NAME" \
    -Fc \
    > "$BACKUP_FILE"

echo "[$(date '+%Y-%m-%d %H:%M:%S')] Backup created: $BACKUP_FILE"

# 30日より古いバックアップを削除
find "$BACKUP_DIR" \
    -type f \
    -name "kaz_app_*.dump" \
    -mtime +30 \
    -delete

echo "[$(date '+%Y-%m-%d %H:%M:%S')] Old backups deleted."