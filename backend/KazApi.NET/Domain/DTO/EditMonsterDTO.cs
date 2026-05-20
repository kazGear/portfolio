using KazApi.Common._Filter;
using System.Text.Json.Serialization;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// モンスターパラメータクラス
    /// </summary>
    public class EditMonsterDTO
    {
        private string _monsterId;
        private string _monsterName;
        private int _monsterType;
        private string _monsterTypeName;
        private int _hp;
        private int _attack;
        private int _speed;
        private int _week;
        private string _weekName;

        [JsonPropertyName("MonsterId")]
        public string MonsterId
        {
            get { return _monsterId; }
            set { _monsterId = Validation.Id(value); }
        }

        [JsonPropertyName("MonsterName")]
        public string MonsterName
        {
            get { return _monsterName; }
            set { _monsterName = Validation.Name(value); }
        }

        [JsonPropertyName("MonsterType")]
        public int MonsterType
        {
            get { return _monsterType; }
            set { _monsterType = value; } // TODO バリデーション
        }

        [JsonPropertyName("MonsterTypeName")]
        public string MonsterTypeName
        {
            get { return _monsterTypeName; }
            set { _monsterTypeName = Validation.Name(value); }
        }

        [JsonPropertyName("Hp")]
        public int Hp
        {
            get { return _hp; }
            set { _hp = Validation.Hp(value); }
        }

        [JsonPropertyName("Attack")]
        public int Attack
        {
            get { return _attack; }
            set { _attack = Validation.Strength(value); }
        }

        [JsonPropertyName("Speed")]
        public int Speed
        {
            get { return _speed; }
            set { _speed = Validation.Strength(value); }
        }

        [JsonPropertyName("Week")]
        public int Week
        {
            get { return _week; }
            set { _week = Validation.ElementType(value); }
        }

        [JsonPropertyName("WeekName")]
        public string WeekName
        {
            get { return _weekName; }
            set { _weekName = Validation.Name(value); }
        }

        [JsonPropertyName("IsDisabled")]
        public string IsDisabled { get; set; }

        // 変更後パラメータ

        [JsonPropertyName("AfterName")]
        public string? AfterName { get; set; }
        [JsonPropertyName("AfterHp")]
        public int? AfterHp { get; set; }
        [JsonPropertyName("AfterAttack")]
        public int? AfterAttack { get; set; }
        [JsonPropertyName("AfterSpeed")]
        public int? AfterSpeed { get; set; }
        [JsonPropertyName("AfterWeek")]
        public int? AfterWeek { get; set; }
        [JsonPropertyName("IsChanged")]
        public bool? IsChanged { get; set; }
    }
}
