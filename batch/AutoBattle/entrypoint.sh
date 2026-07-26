#!/bin/sh

set -e

# コンテナの環境変数を /etc/environment に保存
printenv > /etc/environment

# cronをフォアグラウンドで起動
exec cron -f