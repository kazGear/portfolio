using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    public class UserDTO
    {
        private string _loginId;
        private string _loginPass;
        private int _failedLoginCnt;
        private string _dispName;
        private string _dispShortName;
        private int _role;
        private string _roleName;
        private int _cash;
        private int _wins;
        private int _winsGetCash;
        private int _losses;
        private int _bankruptcyCnt;
        private int _lossesLostCash;

        public string LoginId 
        {
            get { return _loginId; }
            set { _loginId = Validation.LoginId(value); }
        }

        public string LoginPass 
        {
            get { return _loginPass; }
            set { _loginPass = Validation.LoginPass(value); } 
        }

        public int FailedLoginCnt 
        {
            get { return _failedLoginCnt; }
            set { _failedLoginCnt = Validation.Count(value); }
        }

        public bool IsLoginDisabled { get; set; }

        public string DispName 
        {
            get { return _dispName; }
            set { _dispName = Validation.Name(value); }
        }

        public string DispShortName 
        {
            get { return _dispShortName; }
            set { _dispShortName = Validation.ShortName(value); }
        }

        public int Role
        {
            get { return _role; }
            set { _role = Validation.CodeValue(value); }
        }

        public string RoleName 
        {
            get { return _roleName; }
            set { _roleName = Validation.Name(value); }
        }

        public int Cash 
        {
            get { return _cash; }
            set { _cash = Validation.Amount(value); }
        }

        public int Wins 
        {
            get { return _wins; }
            set { _wins = Validation.Count(value); }
        }

        public int WinsGetCash 
        {
            get { return _winsGetCash; }
            set { _winsGetCash = Validation.Amount(value); }
        }

        public int Losses 
        {
            get { return _losses; }
            set { _losses = Validation.Count(value); }
        }

        public int BankruptcyCnt 
        {
            get { return _bankruptcyCnt; }
            set { _bankruptcyCnt = Validation.Count(value); }
        }

        public int LossesLostCash 
        {
            get { return _lossesLostCash; }
            set { _lossesLostCash = Validation.Amount(value); }
        }

        public string UserImage { get; set; }

        public string Token { get; set; }
    }
}