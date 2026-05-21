# docker 起動
# 開発（ローカル）

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml up --detach

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml down

# 本番（VPS）

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml up --detach

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml down

# 開発

win側で開発・修正を行う。/mnt/c/repository/portfolio
win側の成果物をwsl側にコピーし、wsl側も最新を保つ。syncコマンド or cron

# C#バッチを叩く

例：
docker exec -it portfolio-batch-autobattle dotnet /src/batch/autoBattle/bin/Debug/net8.0/autoBattle.dll

# C#デバック + docker ※VSのF5で実行しないこと。

api（ASP.NET コンテナ）が Up になっていればOK。

「プロセスにアタッチ」で Docker の dotnet プロセスに接続する。
VSで、デバッグ → プロセスにアタッチ

「接続対象」を Docker に変更
プロセス一ｘ覧にこういうのが出る：dotnet  (KazApi.dll)
これが Docker の中で動いている ASP.NET API。
これを選んで「アタッチ」
エンジン選択は、マネージド（Unix 用 .NET Core）

### tmp memo
1. Docker Desktop の自動起動が OFF
Settings → General →

Start Docker Desktop when you log in が OFF だと起動しない。