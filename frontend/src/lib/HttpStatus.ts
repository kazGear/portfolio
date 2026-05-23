export const HttpStatus = {
    // Informational（情報レスポンス）

    /** 100: 要求を受理し処理を継続できる。 */
    Continue: 100,
    /** 101: プロトコル切り替え要求を受け入れた。 */
    SwitchingProtocols: 101,
    /** 102: 処理中であり最終応答ではない。 */
    Processing: 102,
    /** 103: 本応答前にヘッダーを先行送信する。 */
    EarlyHints: 103,

    // Success（成功レスポンス）

    /** 200: 要求が正常に処理された。 */
    OK: 200,
    /** 201: 新しいリソースが作成された。 */
    Created: 201,
    /** 202: 要求を受理したが処理は未完了。 */
    Accepted: 202,
    /** 203: 返却情報がオリジンサーバーのものではない。 */
    NonAuthoritativeInformation: 203,
    /** 204: 正常だが返す内容は存在しない。 */
    NoContent: 204,
    /** 205: 表示内容をリセットすべきである。 */
    ResetContent: 205,
    /** 206: レスポンスの一部のみ返されている。 */
    PartialContent: 206,
    /** 207: 複数の状態を持つ結果を返す（WebDAV）。 */
    MultiStatus: 207,
    /** 208: すでに報告済みのリソースである（WebDAV）。 */
    AlreadyReported: 208,
    /** 226: GET 以外のメソッドにも内容を返した。 */
    IMUsed: 226,

    // Redirection（リダイレクション）

    /** 300: 複数の遷移先候補が存在する。 */
    MultipleChoices: 300,
    /** 301: リソースが恒久的に移動した。 */
    MovedPermanently: 301,
    /** 302: 一時的に別の場所へ移動している。 */
    Found: 302,
    /** 303: 別の URI を参照すべきである。 */
    SeeOther: 303,
    /** 304: 変更がないためキャッシュを利用できる。 */
    NotModified: 304,
    /** 305: プロキシ経由でアクセスすべき（非推奨）。 */
    UseProxy: 305,
    /** 307: 一時的に別 URI へリダイレクトされている。 */
    TemporaryRedirect: 307,
    /** 308: 恒久的に別 URI へリダイレクトされている。 */
    PermanentRedirect: 308,

    // Client Error（クライアントエラー）

    /** 400: リクエストが不正で処理できない。 */
    BadRequest: 400,
    /** 401: 認証が必要だが情報が無いか無効。 */
    Unauthorized: 401,
    /** 402: 支払いが必要な状態（予約コード）。 */
    PaymentRequired: 402,
    /** 403: アクセス権限がなく禁止されている。 */
    Forbidden: 403,
    /** 404: 指定されたリソースが見つからない。 */
    NotFound: 404,
    /** 405: 許可されていないメソッドで要求された。 */
    MethodNotAllowed: 405,
    /** 406: 要求内容がクライアントの条件に合わない。 */
    NotAcceptable: 406,
    /** 407: プロキシ認証が必要である。 */
    ProxyAuthenticationRequired: 407,
    /** 408: クライアントのタイムアウトにより接続が切断された。 */
    RequestTimeout: 408,
    /** 409: リソースの状態が競合している。 */
    Conflict: 409,
    /** 410: リソースが削除され利用できない。 */
    Gone: 410,
    /** 411: 必須ヘッダーが不足している。 */
    LengthRequired: 411,
    /** 412: 事前条件ヘッダーが満たされなかった。 */
    PreconditionFailed: 412,
    /** 413: リクエストボディが大きすぎる。 */
    PayloadTooLarge: 413,
    /** 414: URI が長すぎて処理できない。 */
    URITooLong: 414,
    /** 415: サポートされていないメディアタイプ。 */
    UnsupportedMediaType: 415,
    /** 416: 要求範囲がリソース外である。 */
    RangeNotSatisfiable: 416,
    /** 417: Expect ヘッダーの条件が満たされない。 */
    ExpectationFailed: 417,
    /** 418: ティーポットなのでコーヒーは淹れられない（ジョーク）。 */
    ImATeapot: 418,
    /** 421: 不適切なサーバーに要求が送られた。 */
    MisdirectedRequest: 421,
    /** 422: 内容が意味的に不正で処理できない。 */
    UnprocessableEntity: 422,
    /** 423: リソースがロックされている。 */
    Locked: 423,
    /** 424: 依存する要求が失敗したため処理不可。 */
    FailedDependency: 424,
    /** 425: 早すぎるタイミングで送信された。 */
    TooEarly: 425,
    /** 426: プロトコルのアップグレードが必要。 */
    UpgradeRequired: 426,
    /** 428: 事前条件ヘッダーが必須だが指定されていない。 */
    PreconditionRequired: 428,
    /** 429: 一定時間内の要求回数が多すぎる。 */
    TooManyRequests: 429,
    /** 431: リクエストヘッダーが大きすぎる。 */
    RequestHeaderFieldsTooLarge: 431,
    /** 451: 法的理由により利用できない。 */
    UnavailableForLegalReasons: 451,

    // Server Error（サーバエラー）

    /** 500: サーバー内部で予期しないエラーが発生。 */
    InternalServerError: 500,
    /** 501: 要求された機能が実装されていない。 */
    NotImplemented: 501,
    /** 502: 上流サーバーから不正な応答を受け取った。 */
    BadGateway: 502,
    /** 503: サービスが一時的に利用不能。 */
    ServiceUnavailable: 503,
    /** 504: 上流サーバーの応答がタイムアウトした。 */
    GatewayTimeout: 504,
    /** 505: サポートされていない HTTP バージョン。 */
    HTTPVersionNotSupported: 505,
    /** 506: ネゴシエーション設定が不正。 */
    VariantAlsoNegotiates: 506,
    /** 507: サーバーのストレージ容量不足。 */
    InsufficientStorage: 507,
    /** 508: リソース参照がループしている。 */
    LoopDetected: 508,
    /** 510: 追加の拡張が必要だが定義されていない。 */
    NotExtended: 510,
    /** 511: ネットワーク認証が必要。 */
    NetworkAuthenticationRequired: 511,
} as const;
