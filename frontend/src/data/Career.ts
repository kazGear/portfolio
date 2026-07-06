/***********************************************************
 * JsonData, ArrayData
 ************************************************************/

/**
 * メタデータ
 */
export const metaData = {
    "lastUpdate": "2026/07/06",
}

/**
 * 基本情報
 */
export const profile = {
    "myName": "K.O",
    "myAddress": "福岡県朝倉市在住　最寄り駅：甘木（西鉄甘木線）"
}

/**
 * スキルデータ
 */
export const skills = [
    "HTML",
    "CSS",
    "JavaScript(jQuery)",
    "TypeScript",
    "React(Vite)",
    "C#(ASP.NET Core)",
    "Java(Spring Boot)",
    "Go言語",
    "PostgreSQL",
    "SQL Server",
    "My SQL",
    "VBA",
    "基本設計",
    "詳細設計",
    "正規表現",
    "プチAI駆動開発",
    "リファクタリング",
    "リーダブルコーディング",
    "レイヤードアーキテクチャ",
    "テーブルドリブンテスト",
    "【 得意言語：java, C#, SQL 】"
];

/**
 * スキルデータ(習熟度低め・練習中・関心事)
 */
export const littleSkills = [
    "クリーンアーキテクチャ",
    "テーブル設計",
    "SQLチューニング",
    "docker",
    "Linux(ubunts)"
];

/**
 * 使用ツール
 */
export const useTools = [
    "Windows",
    "Cursor",
    "WinMerge",
    "サクラエディタ",
    "Eclipse",
    "Visual Studio",
    "Visual Studio Code",
    "A5Mk2",
    "DBeaver"
];

/**
 * 読んだ本
 */
export const readBooks = [
    "Go言語 プログラミングエッセンス",
    "一流プログラマが教える Go言語 大全",
    "SQLパズル",
    "達人に学ぶSQL徹底指南書",
    "SQLアンチパターン",
    "システム開発・刷新のためのデータモデル大全",
    "Java言語で学ぶデザインパターン入門",
    "Head First デザインパターン",
    "リーダブルコード",
    "達人プログラマー",
    "ルールズ・オブ・プログラミング",
    "プリンシプル オブ プログラミング",
    "実務で役立つシステム設計の原則",
    "good code, bad code",
    "問題解決のための「アルゴリズム × 数学」がしっかり身につく本",
    "リファクタリング",
    "実践ドメイン駆動設計",
    "クリーンアーキテクチャ",
    "ソフトウェア受託現場の「失敗」集めてみた",
    "ソフトウェア開発現場の「失敗」集めてみた",
    "React 開発入門",
    "React 実践の教科書",
    "TypeScriptとReact/Next.js 実践Webアプリケーション開発",
    "安全なWebアプリケーションの作り方",
    "ソフトウェアテストの教科書",
    "始めての自働テスト",
    "ストリートコーダー"
];

/**
 * ポートフォリオ
 */
export const portfolio = {
    "appURLCaption": "成果物・webで稼働中のアプリケーション",
    "appURL": "https://kazapp-trial.com",
    "sourceURLCaption": "ソース（React, TypeScript, C#/ASP.NET CORE, Go言語, PostgreSQL, docker）",
    "sourceURL": "https://github.com/kazGear/portfolio",
    "comment": [
        "[ メインコンテンツ ]",
        "◆ Guitar Gallery（観賞用、自動検索）",
        "◆ Guitar REST API（ギターデータ取得API）",
        "◆ Monster Battle Arena（とある複数のミニゲームをブラウザ上で表現）",
        "◆ web版経歴書",
        "<br/>",
        "React、TypeScript、C#、Go、PostgreSQLを用いて個人開発しているWebアプリケーションです。",
        "バックエンドAPI、データベース設計、フロントエンドまで一人で設計・実装しています。",
        "保守性・拡張性を意識しながら継続的にリファクタリングを行っています。",
        "※chrome, edge推奨。iPad横画面対応",
        "<br/>",
        "◆ Guitar Gallery ◆",
        "複数メーカーのギター情報を横断検索できるWebサービスです。",
        "スクレイピングによるデータ収集基盤を構築し、それによる収集物がギターのデータソースとなります。",
        "検索条件を設定するだけで自動的にギターを検索します。検索ボタンを押す手間などかかりません。",
        "ページネーションやソートにも対応しています。",
        "検索UIや詳細モーダルなど、ユーザビリティも重視した設計を行っています。",
        "<br/>",
        "◆ Monster Battle Arena ◆",
        "どのモンスターが勝ち残るかを賭けるゲームです。モンスター達は自動で戦います。",
        "複数のモンスターたち（約８０種）がランダムで登場し、様々なスキル（約７０種）を駆使して戦います（１モンスターにつき６スキル保持）。",
        "掛け金（ゲーム内通貨）の配当倍率は、モンスターの数、組み合わせで動的に変動します。",
        "モンスター達のステータス、スキルは設定画面から変更することも可能です。",
        "<br/>",
        "担当範囲： 要件定義、設計、実装、スクレイピング、データベース設計、API開発、フロントエンド、デプロイまで全て担当。",
    ]
}

/**
 * 得意分野
 */
export const specialty = [
    "一貫したフルスタック寄りの設計、開発",
    "Java（Spring Boot）、C#（ASP.NET Core）を用いたWebアプリケーション開発",
    "React、TypeScriptを用いたSPA開発（レスポンシブ）",
    "PostgreSQL、SQL Server、MySQLを用いたデータベース設計・SQL実装",
    "複雑な検索・集計・ランキングなどの業務ロジック実装",
    "保守性・可読性・拡張性を意識した設計・リファクタリング",
    "スクレイピング、データ加工、データベース登録までを含むデータ収集基盤の構築",
    "要件整理から設計・実装・テストまで一貫した開発"
];

/**
 * PR事項
 */
export const prPoint = [
    {
        "point": "設計・保守性・品質を意識した開発",
        "comment":
            "機能を実装するだけではなく、将来的な保守や機能追加まで考慮した実装を心掛けています。<br/>" +
            "命名、責務分割、リファクタリングを意識し、可読性・保守性・拡張性の高いコードを書くことを大切にしています。<br/><br/>" +
            "例：<br/>" +
            "・フロントエンド、バックエンド、リポジトリの分割・疎結合化<br/>" +
            "・クラス設計等<br/>" +
            "・抽象化<br/>" +
            "・仕組化<br/>" +
            "・責務の意識<br/>" +
            "・はまる箇所にはデザインパターンの適用<br/>"
    },
    {
        "point": "問題解決能力",
        "comment":
            "あるプロジェクトでは、参画中は延々とバグ潰しをしておりました。個人開発でもバグはでるもので、都度自身で対応しています。<br/>" +
            "論理的に問題を切り分け、仮説を検証し、課題を解決することに努めます。"
    },
    {
        "point": "学習・キャッチアップ能力",
        "comment":
            "未経験の技術についても、調査・検証を行いながら短期間で習得し、実務へ適用してきました。<br/>" +
            "業務では急な業務効率化ツールの開発依頼があり、未習得のPowerShellを技術として指定されましたが、無事乗り切りました。<br/>" +
            "個人開発ではReact、TypeScript、Go、Dockerなど継続的に新しい技術を学び、設計から実装まで取り組んでいます。<br/><br/>" +
            "定期的に技術書を購入し、実践し、スキル向上に努めています。"
    },
    {
        "point": "バックエンドを軸とした一貫した開発経験",
        "comment":
            "Java（Spring Boot）、C#（ASP.NET Core）を中心に、WebUI、API、SQL、バッチ処理まで幅広く担当してきました。<br/>" +
            "フロントエンドだけ、バックエンドだけではなく、システム全体を見ながら実装できることを強みとしています。"
    },
    {
        "point": "データベース・SQLを得意分野としています",
        "comment":
            "PostgreSQL、SQL Server、MySQLを使用した開発経験があります。<br/>" +
            "複雑な検索、集計、ランキングなどのSQL実装も担当しており、パフォーマンスや保守性も考慮しながら実装を行っています。"
    },
    {
        "point": "個人開発",
        "comment":
            "React・TypeScript・C#・Go・PostgreSQLを用いたWebサービスを継続的に開発しています。<br/>" +
            "ブラウザゲーム、認証機能、検索機能、管理機能、ランキング・集計機能などを実装するとともに、スクレイピングによるデータ収集基盤も構築しています。<br/>" +
            "フロントエンド（レスポンシブ）・バックエンド・データベースまで一人で設計・実装を行い、継続的に改善を進めています。"
    },
    {
        "point": "コミュニケーション",
        "comment":
            "腰が低いので接しやすいと思います。<br/>" +
            "技術的背景を踏まえて状況を整理し、必要な情報を簡潔に共有することを意識しています。<br/>" +
            "仕様や要件の意図を確認しながら進めるため、認識齟齬を減らします。"
    }
]

/**
 * 職務経歴データ
 *
 * template
{
    "historyTitle": "...",
    "period": "xxxx.xx ~ xxxx.xx",
    "industry": "...",
    "scale": "...",
    "programmingLanguages": "...",
    "jobContents": [
        内容１，
        内容２，
        ...
    ]
}
 */
export const career = [
    {
        "historyTitle": "基幹システム（リプレイス）",
        "period": "2025.04 ~ 2025.06",
        "industry": "美容",
        "scale": "7名",
        "programmingLanguages": "html, javaScript, jQuery, C#, ASP.NET Core, mySql",
        "jobContents": [
            "・障害対応",
            "・バグ対応チーム配属。",
            "・対応バグ件数: 25件（件数としては1位～2位をキープ）"
        ]
    },
    {
        "historyTitle": "情報系アプリ（新規機能開発）",
        "period": "2025.01 ~ 2025.03",
        "industry": "金融",
        "scale": "15 ~ 20名",
        "programmingLanguages": "html, css, javaScript, jQuery, java, springBoot, postgreSQL",
        "jobContents": [
            "・詳細設計～単体テスト",
            "・新規機能の開発を担当。",
            "・要件等が複雑であったため、リファクタリングしながら丁寧にコーディング。",
            "・製造範囲：WebUI、サーバ処理、SQL",
            "・2022.01 ~ 2022.08 , 2023.01 ~ 2024.03 の期間と同じ現場で3度目の現場入り。"
        ]
    },
    {
        "historyTitle": "基幹システム（総合テスト～）",
        "period": "2024.06 ~ 2024.12",
        "industry": "住宅関連サービス",
        "scale": "30人",
        "programmingLanguages": "C#, ASP.NET, cshtml, javaScript, SQL Server",
        "jobContents": [
            "総合テスト、障害調査、障害対応",
            "・障害発見件数：16件",
            "・障害対応件数：85件以上 ※コーディング、単体テスト",
            "・元々は2カ月で終了予定。"
        ]
    },
    {
        "historyTitle": "帳票出力システム（新規開発）",
        "period": "2024.04 ~ 2024.05",
        "industry": "公共",
        "scale": "2名",
        "programmingLanguages": "C#, ASP.NET, ActiveReport, SQL Server",
        "jobContents": [
            "・詳細設計、コーディング、単体テスト",
            "・5つの帳票出力機能を担当。",
            "・製造範囲：帳票定義ライブラリ、サーバ処理、SQL",
        ]
    },
    {
        "historyTitle": "情報系アプリ（新規機能開発、機能改修、保守運用）",
        "period": "2022.01 ~ 2022.08 , 2023.01 ~ 2024.03",
        "industry": "金融",
        "scale": "15 ~ 20名",
        "programmingLanguages": "html, css, javaScript, jQuery, java, springBoot, postgreSQL",
        "jobContents": [
            "・基本設計～結合テスト、保守運用、調査、ルーティンの自動化（VBA）、新人の方のコーディング指導",
            "・新規機能の製造は4画面、3帳票、バッチ5本を担当。",
            "・保守運用：アプリの障害対応、既存不良の解消、顧客依頼の修正作業等。",
            "・製造範囲：WebUI、サーバ・バッチ処理、SQL"
        ]
    },
    {
        "historyTitle": "経理システム（リプレイス）",
        "period": "2022.11 ~ 2022.12",
        "industry": "電力",
        "scale": "30名",
        "programmingLanguages": "java, springBoot, abap（SAP）",
        "jobContents": [
            "・詳細設計、コーディング",
            "・リプレイス：abap言語（SAP）の仕様を調べながらjavaへ翻訳。",
            "・コーディング：2機能担当（10,000 Step）。",
            "・非効率・旧時代的な社風が合わず自ら退場。"
        ]
    },
    {
        "historyTitle": "情報システム（リプレイス）",
        "period": "2022.09 ~ 2022.10",
        "industry": "公共",
        "scale": "10名",
        "programmingLanguages": "html, css, java, springBoot, struts, mySQL, VBA, PowerShell",
        "jobContents": [
            "・調査、ツール製造、結合テスト",
            "・ツール製造：VBA, powerShell を使用し、顧客の業務効率を改善した。",
            "・リプレイス：フレームワークの置き換え（struts to spring）。",
            "・調査：フレームを置き換えた影響により障害が発生しており、原因を調査。",
            "・エンド様の予算縮小により退場（プロジェクトの急な中止）。"
        ]
    },
    {
        "historyTitle": "WEBアプリ & 基幹アプリ（新規立ち上げ、既存アプリの改修・障害対応）",
        "period": "2021.06 ~ 2021.12",
        "industry": "製造",
        "scale": "1名",
        "programmingLanguages": "html, css, javaScript, C#, VB.net, ASP.net Core, postgreSQL",
        "jobContents": [
            "・要求分析～結合テスト、保守",
            "・VB.NETで作成された既存アプリに対し、機能の追加や改修を行った。",
            "・WEBアプリ（案件管理用）を新規に作成。WebUI、サーバ処理、DB処理、サーバの構築などを行った。",
            "・初現場。顧客側はIT技術者0人 + 単独での常駐であるためプロとしての経験とは言えない。"
        ]
    },
    {
        "historyTitle": "職業訓練（Webアプリ制作）",
        "period": "2020.12 ~ 2021.04",
        "industry": "―",
        "scale": "チーム：4名、クラス：15名",
        "programmingLanguages": "java, postgreSQL, HTML, Servlet, JSP",
        "jobContents": [
            "・独学でプログラムに触れた後、基礎固めのため入校。",
            "・java, postgreSQL, HTML, Servlet, JSP の基本を学習。",
            "・最終課題のチームによるアプリ制作では、詳細設計～実装を担当した。",
            "・実装についてはほぼ1人で担う。",
            "・帰宅後、休日もひたすらコーディングに打ち込んでいた。"
        ]
    },
    {
        "historyTitle": "IT以外の主な経歴",
        "period": "2017.05 ~ 2020.09",
        "industry": "税理士業、経理",
        "scale": "7～40名",
        "programmingLanguages": "VBA",
        "jobContents": [
            "税務会計、決算・申告、監査、経理作業、巡回訪問",
            "VBAを使用した業務効率化（業後、休暇を利用し在職中に習得）"
        ]
    }
];