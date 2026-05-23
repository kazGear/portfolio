namespace CSLib.Lib
{
    public class HttpStatus
    {
        // Informational（情報レスポンス）

        /// <summary>100: 要求を受理し処理を継続できる。</summary>
        public static readonly int Continue = 100;
        /// <summary>101: プロトコル切り替え要求を受け入れた。</summary>
        public static readonly int SwitchingProtocols = 101;
        /// <summary>102: 処理中であり最終応答ではない。</summary>
        public static readonly int Processing = 102;
        /// <summary>103: 本応答前にヘッダーを先行送信する。</summary>
        public static readonly int EarlyHints = 103;

        // Success（成功レスポンス）

        /// <summary>200: 要求が正常に処理された。</summary>
        public static readonly int OK = 200;
        /// <summary>201: 新しいリソースが作成された。</summary>
        public static readonly int Created = 201;
        /// <summary>202: 要求を受理したが処理は未完了。</summary>
        public static readonly int Accepted = 202;
        /// <summary>203: 返却情報がオリジンサーバーのものではない。</summary>
        public static readonly int NonAuthoritativeInformation = 203;
        /// <summary>204: 正常だが返す内容は存在しない。</summary>
        public static readonly int NoContent = 204;
        /// <summary>205: 表示内容をリセットすべきである。</summary>
        public static readonly int ResetContent = 205;
        /// <summary>206: レスポンスの一部のみ返されている。</summary>
        public static readonly int PartialContent = 206;
        /// <summary>207: 複数の状態を持つ結果を返す（WebDAV）。</summary>
        public static readonly int MultiStatus = 207;
        /// <summary>208: すでに報告済みのリソースである（WebDAV）。</summary>
        public static readonly int AlreadyReported = 208;
        /// <summary>226: GET 以外のメソッドにも内容を返した。</summary>
        public static readonly int IMUsed = 226;

        // Redirection（リダイレクション

        /// <summary>300: 複数の遷移先候補が存在する。</summary>
        public static readonly int MultipleChoices = 300;
        /// <summary>301: リソースが恒久的に移動した。</summary>
        public static readonly int MovedPermanently = 301;
        /// <summary>302: 一時的に別の場所へ移動している。</summary>
        public static readonly int Found = 302;
        /// <summary>303: 別の URI を参照すべきである。</summary>
        public static readonly int SeeOther = 303;
        /// <summary>304: 変更がないためキャッシュを利用できる。</summary>
        public static readonly int NotModified = 304;
        /// <summary>305: プロキシ経由でアクセスすべき（非推奨）。</summary>
        public static readonly int UseProxy = 305;
        /// <summary>307: 一時的に別 URI へリダイレクトされている。</summary>
        public static readonly int TemporaryRedirect = 307;
        /// <summary>308: 恒久的に別 URI へリダイレクトされている。</summary>
        public static readonly int PermanentRedirect = 308;

        // Client Error（クライアントエラー）

        /// <summary>400: リクエストが不正で処理できない。</summary>
        public static readonly int BadRequest = 400;
        /// <summary>401: 認証が必要だが情報が無いか無効。</summary>
        public static readonly int Unauthorized = 401;
        /// <summary>402: 支払いが必要な状態（予約コード）。</summary>
        public static readonly int PaymentRequired = 402;
        /// <summary>403: アクセス権限がなく禁止されている。</summary>
        public static readonly int Forbidden = 403;
        /// <summary>404: 指定されたリソースが見つからない。</summary>
        public static readonly int NotFound = 404;
        /// <summary>405: 許可されていないメソッドで要求された。</summary>
        public static readonly int MethodNotAllowed = 405;
        /// <summary>406: 要求内容がクライアントの条件に合わない。</summary>
        public static readonly int NotAcceptable = 406;
        /// <summary>407: プロキシ認証が必要である。</summary>
        public static readonly int ProxyAuthenticationRequired = 407;
        /// <summary>408: クライアントのタイムアウトにより接続が切断された。</summary>
        public static readonly int RequestTimeout = 408;
        /// <summary>409: リソースの状態が競合している。</summary>
        public static readonly int Conflict = 409;
        /// <summary>410: リソースが削除され利用できない。</summary>
        public static readonly int Gone = 410;
        /// <summary>411: 必須ヘッダーが不足している。</summary>
        public static readonly int LengthRequired = 411;
        /// <summary>412: 事前条件ヘッダーが満たされなかった。</summary>
        public static readonly int PreconditionFailed = 412;
        /// <summary>413: リクエストボディが大きすぎる。</summary>
        public static readonly int PayloadTooLarge = 413;
        /// <summary>414: URI が長すぎて処理できない。</summary>
        public static readonly int URITooLong = 414;
        /// <summary>415: サポートされていないメディアタイプ。</summary>
        public static readonly int UnsupportedMediaType = 415;
        /// <summary>416: 要求範囲がリソース外である。</summary>
        public static readonly int RangeNotSatisfiable = 416;
        /// <summary>417: Expect ヘッダーの条件が満たされない。</summary>
        public static readonly int ExpectationFailed = 417;
        /// <summary>418: ティーポットなのでコーヒーは淹れられない（ジョーク）。</summary>
        public static readonly int ImATeapot = 418;
        /// <summary>421: 不適切なサーバーに要求が送られた。</summary>
        public static readonly int MisdirectedRequest = 421;
        /// <summary>422: 内容が意味的に不正で処理できない。</summary>
        public static readonly int UnprocessableEntity = 422;
        /// <summary>423: リソースがロックされている。</summary>
        public static readonly int Locked = 423;
        /// <summary>424: 依存する要求が失敗したため処理不可。</summary>
        public static readonly int FailedDependency = 424;
        /// <summary>425: 早すぎるタイミングで送信された。</summary>
        public static readonly int TooEarly = 425;
        /// <summary>426: プロトコルのアップグレードが必要。</summary>
        public static readonly int UpgradeRequired = 426;
        /// <summary>428: 事前条件ヘッダーが必須だが指定されていない。</summary>
        public static readonly int PreconditionRequired = 428;
        /// <summary>429: 一定時間内の要求回数が多すぎる。</summary>
        public static readonly int TooManyRequests = 429;
        /// <summary>431: リクエストヘッダーが大きすぎる。</summary>
        public static readonly int RequestHeaderFieldsTooLarge = 431;
        /// <summary>451: 法的理由により利用できない。</summary>
        public static readonly int UnavailableForLegalReasons = 451;

        // Server Error（サーバエラー）

        /// <summary>500: サーバー内部で予期しないエラーが発生。</summary>
        public static readonly int InternalServerError = 500;
        /// <summary>501: 要求された機能が実装されていない。</summary>
        public static readonly int NotImplemented = 501;
        /// <summary>502: 上流サーバーから不正な応答を受け取った。</summary>
        public static readonly int BadGateway = 502;
        /// <summary>503: サービスが一時的に利用不能。</summary>
        public static readonly int ServiceUnavailable = 503;
        /// <summary>504: 上流サーバーの応答がタイムアウトした。</summary>
        public static readonly int GatewayTimeout = 504;
        /// <summary>505: サポートされていない HTTP バージョン。</summary>
        public static readonly int HTTPVersionNotSupported = 505;
        /// <summary>506: ネゴシエーション設定が不正。</summary>
        public static readonly int VariantAlsoNegotiates = 506;
        /// <summary>507: サーバーのストレージ容量不足。</summary>
        public static readonly int InsufficientStorage = 507;
        /// <summary>508: リソース参照がループしている。</summary>
        public static readonly int LoopDetected = 508;
        /// <summary>510: 追加の拡張が必要だが定義されていない。</summary>
        public static readonly int NotExtended = 510;
        /// <summary>511: ネットワーク認証が必要。</summary>
        public static readonly int NetworkAuthenticationRequired = 511;
    }
}