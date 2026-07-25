#!/bin/sh

set -e

# コンテナの環境変数を /etc/environment に保存
printenv > /etc/environment

# cronに登録する
crontab /etc/cron.d/app-cron

# cronをフォアグラウンドで起動
exec cron -f