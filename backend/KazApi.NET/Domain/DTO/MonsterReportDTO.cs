using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    public class MonsterReportDTO
    {
        private string _monsterId;
        private string _monsterName;
        private int _battleCount;
        private int _wins;

        public string MonsterId
        {
            get { return _monsterId; }
            set {  _monsterId = Validation.Id(value); }
        }

        public string MonsterName
        { 
            get { return _monsterName; }
            set { _monsterName = Validation.Name(value);}
        }

        public int BattleCount
        {
            get { return _battleCount; }
            set { _battleCount = Validation.Count(value); }
        }

        public int Wins
        {
            get { return _wins; }
            set { _wins = Validation.Count(value); }
        }

        public string WinRate { get; set; }
    }
}
