using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    public class BattleReportDTO
    {
        private int _serial;
        private string _monsterId;
        private string _monsterName;

        public int BattleId { get; set; }
        public DateTime BattleEndDate { get; set; }
        public string BattleEndDateStr { get; set; }
        public TimeSpan BattleEndTime { get; set; }
        public string BattleEndTimeStr { get; set; }
        public bool IsWin { get; set; }

        public int Serial
        {
            get { return _serial; }
            set { _serial = Validation.BattleReportSerial(value); }
        }

        public string MonsterId
        {
            get { return _monsterId; }
            set { _monsterId = Validation.Id(value); }
        }

        public string MonsterName
        {
            get { return _monsterName; }
            set { _monsterName = Validation.Name(value); }
        }
    }
}
