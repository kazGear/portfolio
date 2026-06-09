# docker 起動
# 開発（ローカル）

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml up --detach

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml down

# 本番（VPS）

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml up --detach

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml down

# 開発

win側で行う。動作確認、リリース前の確認は,wsl + docker
git管理はC:/repository/portfolio

# C#バッチを叩く

swaggerから実行

# C#デバック + wsl + docker ※VSのF5で実行しないこと。

api（ASP.NET コンテナ）が Up になっていればOK。

「プロセスにアタッチ」で Docker の dotnet プロセスに接続する。
VSで、デバッグ → プロセスにアタッチ

「接続対象」を Docker に変更
プロセス一ｘ覧にこういうのが出る：dotnet  (KazApi.dll)
これが Docker の中で動いている ASP.NET API。
これを選んで「アタッチ」
エンジン選択は、マネージド（Unix 用 .NET Core）

# webCrawler

処理の流れは基本的に抽象化してあるため、具体的な処理を追加したい場合、具体的な処理をするファイルを追加すること。フレーム部分は基本的に修正する必要はない。
例：
fenderのギターをスクレイプしたい >>> scraper_guitar_fender.go を追加する。
gibsonのギターをスクレイプしたい >>> scraper_guitar_gibson.go を追加する。

scraper_guitar.go, scraper.goなど抽象的なファイル名には、抽象化処理、基盤的な処理といった。フレーム部分になる。

・SPA判定基準
1. ページソース（Ctrl+U）にデータがあるか？ 注：初手で分析するのが重要
スペックがある、画像URLがある、テキストがある
→ colly でOK

2. Elements（F12）にだけデータがあるか？
ソースには無いのに Elements にはある
→ JS で生成 → chromedp 必須

3. XHR の中に “データ取得系” があるか？
/api/... /json/... /graphql
→ SPA or 動的 → chromedp