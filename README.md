# docker 起動
# 開発（ローカル）

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml up --detach

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml down

# 本番（VPS）

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml up --detach

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml down

# 開発

wsl側で開発・修正を行う。
wsl側の成果物をwin側にコピーし、win側も最新を保つ。syncコマンド or cron
git管理はC:/repository/portfolio

# C#バッチを叩く

例：
・docker exec -it portfolio-batch-autobattle dotnet /src/batch/autoBattle/bin/Debug/net8.0/autoBattle.dll
・swaggerから実行

# C#デバック + docker ※VSのF5で実行しないこと。

api（ASP.NET コンテナ）が Up になっていればOK。

「プロセスにアタッチ」で Docker の dotnet プロセスに接続する。
VSで、デバッグ → プロセスにアタッチ

「接続対象」を Docker に変更
プロセス一ｘ覧にこういうのが出る：dotnet  (KazApi.dll)
これが Docker の中で動いている ASP.NET API。
これを選んで「アタッチ」
エンジン選択は、マネージド（Unix 用 .NET Core）
