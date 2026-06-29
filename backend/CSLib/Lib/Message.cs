namespace CSLib.Lib
{
    /// <summary>
    /// レスポンス用DTO
    /// </summary>
    public class ApiMessage
    {
        public string? Message { get; set; }
    }
    /// <summary>
    /// api response message
    /// </summary>
    public class Message
    {
        /// <summary>
        /// エラーメッセージ
        /// </summary>
        public static ApiMessage Create(Exception e)
        {
            Console.WriteLine(e.ToString()); // 開発用ヒント
            return new ApiMessage { Message = e.Message }; // API response
        }
        /// <summary>
        /// カスタムエラーメッセージ
        /// </summary>
        public static ApiMessage Create(Exception e, string message)
        {
            Console.WriteLine(e.ToString()); // 開発用ヒント
            return new ApiMessage { Message = $"{message}\n{e.Message}" }; // API response
        }
        /// <summary>
        /// 通常メッセージ
        /// </summary>
        public static ApiMessage Create(string message)
        {
            return new ApiMessage { Message = message };
        }
    }
}
