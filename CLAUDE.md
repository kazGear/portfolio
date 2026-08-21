# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## リポジトリ概要

複数のWebサービス（Monster Battle Arena / Guitar Gallery / web版経歴書 / 案件検索・分析）を
1つのプラットフォームへ統合するモノレポ。4つのランタイムが同居する。

| ディレクトリ | 技術 | 役割 |
| --- | --- | --- |
| `frontend/` | React 18 + Vite + TypeScript + styled-components | SPA（全サービス共通のUI） |
| `backend/PublicApi/` | ASP.NET Core 8 (net8.0) | 認証不要の公開REST API（`/public/v1/...`） |
| `backend/PrivateApi/` | ASP.NET Core 8 (net8.0) | アプリ用API（`/api/...`）。JWT認証あり |
| `backend/CSLib/`, `backend/Repository/` | C# クラスライブラリ | 共通基盤（ログ・通知・例外・定数）とDapper+Npgsqlのデータアクセス |
| `batch/goBatch/` | Go 1.26（モジュール `github.com/kazGear/portfolio/goBatch`） | スクレイピング・バッチ監視・DBアーカイブ |
| `batch/AutoBattle/` | C# コンソール | 毎日の自動バトル実行バッチ（PrivateApiのDomain層を再利用） |
| `infrastructure/` | nginx, PostgreSQL 17 初期化SQL | リバースプロキシ、DB init |

`go.work` はルートから `./batch/goBatch` を参照するためのもの。C#は `backend/KazApi.sln`。

## よく使うコマンド

### Docker（推奨の起動方法）

```bash
# 開発
docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml up --build --detach
docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml ps
docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml down

# 本番（VPS。通常は ./deploy.sh 経由）
docker compose --env-file .env.prod -f compose.base.yaml -f compose.prod.yaml up --build --detach
```

`compose.base.yaml`（ネットワーク・TZ・依存関係）に `compose.dev.yaml` / `compose.prod.yaml` を
必ず重ねて指定する。単体では起動しない。

ログ・切り分け:

```bash
docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml logs --tail=200 <service>
docker exec -it <container> bash
```

バッチの手動実行（cron待ちしない）:

```bash
docker compose --env-file .env.dev -f compose.base.yaml -f compose.dev.yaml exec batch-guitar-crawler ./guitar
```

### frontend

```bash
cd frontend
npm ci
npm run dev      # vite --mode dev  → .env.dev が必要
npm run build    # vite build --mode prod → .env.prod が必要
```

`npm run dev` 単体起動時は `vite.config.ts` の proxy で `/public` → `localhost:7170`、
`/api` → `localhost:5000` に転送される（= dotnet を別途ローカル起動している前提）。
Docker経由の場合はnginxがルーティングするため、`frontend/.env.*` のbase URLは空文字（相対URL）。

lint・テストのスクリプトは未設定。`src/App.test.tsx` / `setupTests.ts` / `reportWebVitals.ts` は
CRA由来の未使用ファイル。

### backend (C#)

```bash
dotnet build backend/KazApi.sln
dotnet run --project backend/PublicApi     # http://localhost:7170 (swagger)
dotnet run --project backend/PrivateApi    # http://localhost:5000
```

テストプロジェクトは存在しない。C#バッチの動作確認はSwaggerから該当APIを叩く運用。

### batch (Go)

```bash
cd batch/goBatch
go build ./...
go test ./...
go test ./internal/crawler/scraper -run TestXxx -v   # 単体テスト指定
go run ./cmd/guitar        # ギタークローラー
go run ./cmd/job           # 案件クローラー
go run ./cmd/batchMonitor  # バッチ異常検知 → Discord通知
go run ./cmd/dbArchiver    # pg_dump によるDBバックアップ
```

テストがあるのは `internal/crawler/scraper/*_test.go` と `pkg/utils/utils_test.go` のみ。

注意: Dockerイメージ（`batch/goBatch/Dockerfile.dev|prod`）は `guitar` / `batchMonitor` /
`dbArchiver` の3つしかビルドしない。`cmd/job` はコンテナに含まれないため、追加するなら
Dockerfile と `cron/*.txt` の両方に追記が必要。

### DB

```powershell
# init.sql の再生成（PowerShell）
pg_dump -U postgres -h localhost -p 5432 -d kaz_app --no-owner --no-privileges -f C:\repository\portfolio\infrastructure\db\init\init.sql

# dumpからのリストア（DBを削除・再作成してから）
pg_restore -h localhost -p 5432 -U postgres -d kaz_app .\infrastructure\db\backup\kaz_app_yyyymmdd_hhmmss.dump
```

`infrastructure/db/init/` はコンテナ初回起動時のみ実行される（`pg-data` ボリュームが空のとき）。

### DBスキーマ

命名は `m_*`（マスタ）/ `t_*`（トランザクション・蓄積データ）。`init.sql` に入っているのは16テーブルで、
モンスターバトル系（`m_monster`, `m_monster_skills`, `m_skill`, `m_item`, `m_shop*`, `m_user`,
`t_battle_result`, `t_my_item`）、コード値マスタ `m_code`、ギター `t_guitars`、
バッチ基盤（`m_batch_config`, `t_batch_execution`）。
`m_monster_origin` / `m_monster_skills_origin` は編集リセット用の初期値。

**案件（Job）系のテーブルは `init.sql` に含まれていない**（`t_jobs`, `t_job_features`,
`t_job_features_created` — Go側が投入、C#側が参照）。空ボリュームから起動したDBでは案件機能が動かないため、
dumpからのリストアか手動DDLが必要。`init.sql` を再生成すれば取り込まれる。

## アーキテクチャ

### リクエストの流れ

```
ブラウザ → nginx(80/443)
             ├ /            → frontend:5173（dev） / ビルド済み静的ファイル（prod）
             ├ /public/v1/  → public-api:5000
             └ /api/        → private-api:5000
                                    ↓
                            PostgreSQL 17 (service名 db)
```

prodでは `infrastructure/nginx/Dockerfile.prod` が React を multi-stage build して nginx に載せるため、
prodのcomposeに frontend サービスは存在しない。HTTPSはnginxで終端し、certbotコンテナが12時間ごとに更新。

### C# 3層 + SQL分離

`Controller → Service → IDatabase(PostgreSQL/Dapper) → sql/*SQL.cs`

- SQL文は `backend/Repository/Repository/sql/` の `static class XxxSQL` に文字列として集約。
  検索条件は `CreateConditions()` などでSQL断片を組み立て、値は必ず Dapper の `DynamicParameters` でバインドする。
- Controller は `new XxxService(configuration)` で直接生成する（Serviceの DI 登録はしていない）。
- 接続文字列は `CSLib/Lib/ConnectionString.Get()` が決定する。
  コンテナ内（`DOTNET_RUNNING_IN_CONTAINER=true`）は `DB_*` 環境変数から組み立て、
  ローカル実行時は `appsettings.{Environment}.json` の `ConnectionStrings:DefaultConnection` を使う。
- 一覧系レスポンスは `TotalCount / Page / PageSize / TotalPages / HasPrev / HasNext + 配列` の形に統一
  （`GuitarsResponse`, `JobsResponse`）。ページ範囲外はエラーではなく空配列。
- PrivateApi は `PropertyNamingPolicy = null` を設定しており、**JSONのキーはC#のPascalCaseそのまま**。
  フロントの型定義（`frontend/src/types/`）もPascalCaseで揃える必要がある。
- 例外は `CSLib/Middleware/ExceptionHandler.cs`（`ExceptionMiddleware`）で一括処理し、
  Discord通知 + Serilogログ + 500レスポンス（`{ message }`）を返す。
  React StrictModeの二重呼び出し対策として `notifyKey`（時刻時単位+パス+例外種別）で重複通知を抑止している。
- PrivateApi のみ `Startup` クラス構成（`Program.cs` 内）。JWT設定・CORS許可オリジンもここ。
  PublicApi は minimal hosting の `Program.cs`。
- コード値（属性・状態・対象種別など）は `m_code` テーブル ＋ `CSLib/Const/C*.cs`（`Enumeration<T>` 派生の
  `static readonly` 定数）の対応で管理。フロントへは `/api/common/FetchElementCode` で配る。

### 認証フロー（既存の実装に合わせること）

`POST /api/auth/login` → `AuthService.AuthenticateUser` → `Jwt.GenerateJwtToken` でトークン発行し、
`UserDTO.Token` として返す → フロントが localStorage（`KEYS.TOKEN`）に保存 →
`apiClient` が全リクエストの `Authorization` ヘッダに**生のトークン（`Bearer ` 接頭辞なし）**を載せる →
有効性確認は `POST /api/auth/checkToken` で `Jwt.IsValidToken()` を自前で呼ぶ（`useHooksOfCommon` /
`useHooksOfIndex` から実行）。

`AddJwtBearer` の設定はしてあるが `[Authorize]` 属性はどこにも付いていないため、
エンドポイントの保護はフレームワークではなく上記の自前チェックで行っている。
`[Authorize]` を後から付ける場合、ヘッダが `Bearer ` 形式でないため `apiClient` 側も直す必要がある。

### Go クローラー（抽象化フレーム + メーカー別実装）

`cmd/<batch>/main.go` → `internal/crawler/service` → `internal/crawler/scraper` → `internal/crawler/repository`

- 共通処理は `scraper_core.go`（並列実行・待機・静的/動的判定）と
  `scraper_core_guitar.go` / `scraper_core_job.go` にある。**フレーム側は基本的に触らない。**
- サイトを追加する場合は `guitar_<maker>.go` / `job_<site>.go` を新規作成し、
  `Scraper[T]` / `PageProvider` / `ModelParser[T]` を実装、
  `service/guitar.go` の `makersFactory()` または `service/job.go` の `jobBoardFactory()` に登録する。
- 静的HTMLは colly、JSレンダリングが必要なページは chromedp を使い分ける。判定手順は `_docs/memo.md` の
  「SPA判定基準」参照。デバッグ時は `service_core.go` の `createChromedpCtxDebug()`（headless=false）に差し替える。
- 抽出値は `map[string]string` に詰めて受け渡す。キーは `pkg/constants`（`C` エイリアスでimport）の
  定数名（`C.Title`, `C.MinSalaryAtMonth` など）を使い、`BuildModel` がそこからモデルを組む。
  項目を増やすときは constants → model → `repository/sql` の3箇所を揃える。
- 並列数は `PARALLEL_COUNT_GUITAR` / `PARALLEL_COUNT_JOB`、ログ保持数は `LOGS_KEEP_COUNT` で制御。
  メーカーごとに `utils.NewLogger` でログファイルを分ける。
- 案件サイトは商品ページのURLが連番のため、`.env` の `PAGE_ID_FROM_*` 〜 `PAGE_ID_TO_*` で
  巡回範囲を指定する方式（サイト追加時はこのペアも追加）。
- 各バッチは開始時に `batchLogger` でDBへ開始ログを入れ、`recover()` でpanicを捕まえて
  エラーステータスを記録する。`cmd/batchMonitor` がそのログを見て未実行・異常をDiscord通知する。

### バッチのスケジュール（cronはコンテナ内、`batch/goBatch/cron/*.txt` を bind mount）

| バッチ | 時刻 | 内容 |
| --- | --- | --- |
| `batch-auto-battle` | 20:01 | モンスター自動バトル10回 + 戦績保存 |
| `batch-db-archiver` | 20:30 | DBダンプを `infrastructure/db/backup/` へ |
| `batch-guitar-crawler` | 21:01 | ギター情報クロール |
| `batch-monitor` | 09:00 | バッチ実行状況の異常検知 → Discord |

cronファイルは**末尾を改行（LF）で終わらせる**こと。CRLFが混ざるとcronが起動しない。
`entrypoint.sh` が `printenv > /etc/environment` してから cron を前面起動する仕組みのため、
環境変数はcron行側で `. /etc/environment` して読み込んでいる。

### frontend

- エントリは `index.html` → `src/main.tsx` → `./App`。**Viteの解決順で `src/App.jsx` が読まれる**
  （`src/App.tsx` はCRA由来の未使用ファイル）。ルーティング追加は `src/App.jsx` の `<Routes>` に対して行う。
- ページは `src/pages/*Page.tsx`、ページ専用の部品は `src/components/<pageName>/`、
  共通部品は `src/components/common/`。クエリパラメータやフェッチのロジックは `src/hooks/useHooksOf*.ts` /
  `useGuitarParams.ts` / `useJobParams.ts` に寄せる。
- API呼び出しは必ず `src/lib/apiClient.ts` の `api.GET/POST/PUT/DELETE` を通す。
  localStorage のトークンを `Authorization` ヘッダに自動付与し、非2xxは `ApiError` を投げる。
  受け側は `useApiErrorHandler`（500系は `/ErrorPage` へ遷移）で処理し、
  最上位は `CommonErrorBoundary` で包む。
- エンドポイントURLは `src/lib/Constants.ts` に集約（`config/env.ts` のbase URLと連結）。色・サイズもここ。
  API を追加したら Controller / `Constants.ts` / `src/types/` の3点セットで更新する。
- スタイルは styled-components。インデントは4スペース（`.prettierrc` / `.vscode/settings.json`）。
- グラフは `@nivo/bar`。案件分析の各グラフは `src/components/JobAnalyzePage/`
  （`ProjectUsageByFeature` / `WorkPlaceByPrefecture` / `SalaryRangeByFeature`）で、
  対応APIは `/api/projectUsageByFeature/get`・`/api/workPlaceByPrefecture/get`・
  `/api/salaryRangeByFeature/get`（いずれも `JobController`）。
  集計は基本SQL側（`JobSQL`）で行い、フロントは描画に寄せる。
  なおこのディレクトリのみPascalCase、他のページ配下は `jobPage/` のようなcamelCase。

## 開発・デプロイフロー

`develop` で作業 → commit → push → GitHubで `develop → main` のPR → merge が**デプロイのトリガー**。
`.github/workflows/deploy.yml` が VPS へSSHして `./deploy.sh` を実行し、
`git pull origin main` → `docker compose up --build -d` → 全コンテナのstatus検証（1つでも非runningならexit 1）。

main に直接pushしない。ローカルのgit管理は `C:/repository/portfolio`、動作確認は WSL + Docker で行う。

## 環境ファイル（すべて .gitignore 対象・コミットしない）

- `.env.dev` / `.env.prod`（ルート）: DB接続、Discord Webhook、クロール対象pageIdの範囲（`PAGE_ID_FROM_*` / `PAGE_ID_TO_*`）、nginxポート
- `frontend/.env.dev` / `.env.prod`: `VITE_PUBLIC_API_BASE_URL` / `VITE_PRIVATE_API_BASE_URL`
- `backend/*/appsettings*.json`: ローカル実行用の接続文字列、`Jwt:Key`、Discord Webhook
- `**/Properties/launchSettings*.json`

環境ファイル・`infrastructure/db/init/init.sql`・`*.dump` は追跡されていないため、
新しい環境では手動配置が必要。

## Docker + Visual Studio でのデバッグ

VSのF5では起動しない。コンテナが Up の状態で
「デバッグ → プロセスにアタッチ」→ 接続対象を Docker → `dotnet (KazApi.dll)` を選択 →
エンジンは「マネージド（Unix 用 .NET Core）」。

## その他

- `_docs/memo.md` に運用手順の一次情報（docker操作、cron、pg_dump、SPA判定など）がまとまっている。
  運用手順を変えたらここも更新する。
- コード内のコメント・ログメッセージは日本語。C#/Goともに `=` の位置を揃える整形スタイルが多用されている。
- ログは `backend/*/logs/`（Serilog、日付ローテート）と
  `batch/goBatch/internal/crawler/logs/<種別>/<サイト名>_<日付>.log`。
