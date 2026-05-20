using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// コードパラメータ for DB
    /// </summary>
    public class CodeDTO
    {
        private string _codeId;
        private int _value;
        private string _name;
        private string _shortName;

        public string CodeId
        {
            get { return _codeId; }
            set { _codeId = Validation.Id(value); }
        }

        public int Value
        {
            get { return _value; }
            set { _value = Validation.CodeValue(value); }
        }

        public string Name
        {
            get { return _name; }
            set { _name = Validation.Name(value); }
        }

        public string ShortName
        {
            get { return _shortName; }
            set { _shortName = Validation.ShortName(value); }
        }

        public string Param1;
        public int Param2;
        public double Param3;
        public string Remarks;
    }
}
