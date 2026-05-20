using System.Reflection;

namespace KazApi.Domain._Const
{
    /// <summary>
    /// 定数値基底クラス
    /// </summary>
    public abstract class Enumeration<T>
    {
        public readonly T Value;
        public readonly string Name;
        public readonly string ShortName;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        protected Enumeration(T value, string name, string shortName = "")
        {
            Value = value;
            Name = name;
            ShortName = shortName;
        }

        public override string ToString() 
            => $"NAME: {Name}, VALUE: {Value}, SHORT_NAME: {ShortName}" ?? "";

        /// <summary>
        /// フィールド名称の列挙を取得
        /// </summary>
        public static IEnumerable<string> FieldNames()
        {
            IList<string> result = [];
            Type fields = typeof(T);

            foreach (FieldInfo field in fields.GetFields())
            {
                result.Add(field.Name);
            }
            return result;
        }

        /// <summary>
        /// フィールド値の列挙を取得
        /// </summary>
        public static IEnumerable<object> FieldValues()
        {
            IList<object> result = [];
            Type fields = typeof(T);
            Type disabled = typeof(Enumeration<T>);

            foreach (FieldInfo field in fields.GetFields())
            {
                try
                {
                    result.Add(field.GetValue(fields) ?? "");
                }
                catch (ArgumentException)
                { }
            }
            return result;
        }
    }
}
