# Guitar data API （説明書）

ギター情報を検索・取得するためのREST APIです。

メーカー、シリーズ、カラー、ボディ材、価格帯などの条件で絞り込み検索ができます。
ページネーションやソートにも対応しています。

特徴 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

メーカー・シリーズ・カラーなど複数条件検索
部分一致検索（name / series）
価格帯検索
ソート機能
ページネーション対応
総件数取得
検索条件なしでも一覧取得可能

エンドポイント -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

GET /api/v1/guitars

クエリパラメータ -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Parameter,Type,Description

makerCd,int,メーカーコード
name,string,ギター名（部分一致検索）
series,string,シリーズ名（部分一致検索）
colorCd,int,カラーコード
bodyMaterialTopCd,int,ボディトップ材
bodyMaterialBackCd,int,ボディバック材
minPrice,int,最低価格
maxPrice,int,最高価格
sort,string,maker / name / price
order,string,ASC / DESC（必須）
page,int,ページ番号（必須）
pageSize,int,1ページの件数（必須）

使用例 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

GET /api/v1/guitars?makerCd=1&series=Strat&page=1&pageSize=25

レスポンス例 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

{
  "totalCount": 283,
  "page": 1,
  "pageSize": 25,
  "totalPages": 12,
  "hasPrev": false,
  "hasNext": true,
  "guitars": [
    {
      "maker": "Fender",
      "name": "American Professional II Stratocaster",
      ...
    },
    {
        ...
    },
  ]
}

検索仕様 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

name, series は部分一致検索を行います。
検索条件を指定しない場合は、全件取得します。

存在しないページを指定した場合はエラーではなく、空のギター配列を返します。


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