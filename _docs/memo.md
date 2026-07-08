# docker 起動
# 開発（ローカル）

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml up --build --detach

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml ps(生存確認用)

docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml down

# 本番（VPS）

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml up --build --detach

docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml ps(生存確認用)

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
→ JS で生成 → chromedp 必須? or networkタブ分析

3. XHR の中に “データ取得系” があるか？
/api/... /json/... /graphql, xhr, fetch
→ SPA or 動的 → chromedp

# react 起動

npm run dev -- --mode dev（開発モード、.env.devが在る前提）
プロジェクトルート（package.json がある場所）で実行。

npm run build -- --mode prod（本番、.env.prodが在る前提）

package.json の script に "dev": "vite --mode dev", "build": "vite build --mode prod" の設定していれば、-- --mode dev 等のコマンド入力は不要。

# README スクショ

htmlで書く。
画像を保存しておき、img src=... で読み込む。

GIFは意外と簡単らしい。
一番おすすめ（Windows）ScreenToGif
無料で神ソフト。

できることは
録画、不要部分を削除、フレーム削除、速度変更、GIF保存

例えば

検索条件を入力
↓
検索ボタン
↓
カード一覧表示
↓
詳細を開く

これを5秒くらい録画するだけ。

# crontab

最後は改行で終了すること。.....\nで終わる必要がある。
分   時   日   月   曜日
0    20   *    *    *    <- 20時に一回実行
* * * * * は1分ごとに実行。

# docker DB 初期化ファイル作成

powerShellで、
pg_dump -U postgres -h localhost -p 5432 -d kaz_app --no-owner --no-privileges -f C:\repository\portfolio\infrastructure\db\init\init.sql

# docker コンテナ内の確認

「Failed to fetch ＝ fetchの問題とは限らない」
「まずAPIが本当に起動しているか確認する」
「docker logs は最初に見る」
「コンテナ内から curl すると切り分けが速い」

コンテナ一覧を見る
docker compose ps
NAME（コンテナ名） を見る。

コンテナに入る
docker exec -it コンテナ名 bash

ログ確認
docker logs コンテナ名

# docker 基本操作

docker compose up -d
docker compose down
docker compose ps
docker compose logs
docker compose exec
docker compose build
docker images
docker volume ls
docker volume prune

# linux 基本操作

ls
cd
pwd
cat
grep
find
tail
less