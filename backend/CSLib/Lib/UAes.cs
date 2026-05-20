using System.Security.Cryptography;
using System.Text;

namespace CSLib.Lib
{
    /// <summary>
    /// 暗号ユーティリティ
    /// </summary>
    public class UAes
    {
        // AES暗号化 key生成するための文字列 (256bitキー(32文字))
        private const string _aesKey = "12345678901234567890123456789012";
        // AES暗号化 初期化ベクトルを生成するための文字列 (128bit(16文字))
        private const string _aesIv = "1234567890123456";

        /// <summary>
        /// コンストラクタ
        /// </summary>
        private UAes()
        {

        }

        /// <summary>
        /// AESで暗号化する関数 
        /// </summary>
        public static string AesEncrypt(string plainText)
        {
            // 暗号化した文字列格納用
            string encrypted_str;

            // Aesオブジェクトを作成
            using (Aes aes = Aes.Create())
            {
                // Encryptorを作成
                using ICryptoTransform encryptor =
                    aes.CreateEncryptor(Encoding.UTF8.GetBytes(_aesKey), Encoding.UTF8.GetBytes(_aesIv));
                // 出力ストリームを作成
                using MemoryStream out_stream = new();
                // 暗号化して書き出す
                using (CryptoStream cs = new(out_stream, encryptor, CryptoStreamMode.Write))
                {
                    using StreamWriter sw = new(cs);
                    // 出力ストリームに書き出し
                    sw.Write(plainText);
                }
                // Base64文字列にする
                byte[] result = out_stream.ToArray();
                encrypted_str = Convert.ToBase64String(result);
            }

            return encrypted_str;
        }

        /// <summary>
        /// AESで復号する関数 
        /// </summary>
        public static string AesDecrypt(string base64text)
        {
            string plain_text;

            // Base64文字列をバイト型配列に変換
            byte[] cipher = Convert.FromBase64String(base64text);

            // AESオブジェクトを作成
            using (Aes aes = Aes.Create())
            {
                // 復号器を作成
                using ICryptoTransform decryptor =
                    aes.CreateDecryptor(Encoding.UTF8.GetBytes(_aesKey), Encoding.UTF8.GetBytes(_aesIv));
                // 復号用ストリームを作成
                using MemoryStream in_stream = new(cipher);
                // 一気に復号
                using CryptoStream cs = new(in_stream, decryptor, CryptoStreamMode.Read);
                using StreamReader sr = new(cs);
                plain_text = sr.ReadToEnd();
            }
            return plain_text;
        }
    }
}
