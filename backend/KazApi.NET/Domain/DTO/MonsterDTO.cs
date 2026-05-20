using KazApi.Common._Filter;
using KazApi.Domain._Monster;
using System.Text.Json.Serialization;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// モンスターパラメータクラス
    /// </summary>
    public class MonsterDTO
    {
        private string _monsterId;
        private string _monsterName;
        private int _monsterType;
        private int _hp;
        private int _maxHp;
        private int _attack;
        private int _defaultAttack;
        private int _speed;
        private int _defaultSpeed;
        private double _dodge;
        private double _defaultDodge;
        private int _week;

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
            set { _monsterType = value; } 
        }
       
        [JsonPropertyName("Hp")]
        public int Hp
        {
            get { return _hp; }
            set { _hp = Validation.Hp(value); }
        }
      
        [JsonPropertyName("MaxHp")]
        public int MaxHp
        {
            get { return _maxHp; }
            set { _maxHp = Validation.Hp(value); }
        }

        [JsonPropertyName("Attack")]
        public int Attack
        {
            get { return _attack; }
            set { _attack = Validation.Strength(value); }
        }

        [JsonPropertyName("DefaultAttack")]
        public int DefaultAttack
        {
            get { return _defaultAttack; }
            set { _defaultAttack = Validation.Strength(value); }
        }

        [JsonPropertyName("Speed")]
        public int Speed
        {
            get { return _speed; }
            set { _speed = Validation.Strength(value); }
        }

        [JsonPropertyName("DefaultSpeed")]
        public int DefaultSpeed
        {
            get { return _defaultSpeed; }
            set { _defaultSpeed = Validation.Strength(value); }
        }

        [JsonPropertyName("Dodge")]
        public double Dodge
        {
            get { return _dodge; }
            set { _dodge = Validation.Rate(value); }
        }

        [JsonPropertyName("DefaultDodge")]
        public double DefaultDodge
        {
            get { return _defaultDodge; }
            set { _defaultDodge = Validation.Rate(value); }
        }

        [JsonPropertyName("Week")]
        public int Week
        {
            get { return _week; }
            set { _week = Validation.ElementType(value); }
        }

        [JsonPropertyName("Skills")]
        public IEnumerable<SkillDTO> Skills { get; set; } = [];

        [JsonPropertyName("Status")]
        public IEnumerable<StateDTO> Status { get; set; } = [];

        [JsonPropertyName("BetScore")]
        public double BetScore { get; set; }

        [JsonPropertyName("BetRate")]
        public double BetRate { get; set; }

        /// <summary>
        /// コンストラクタ
        /// デシリアライズのため必須
        /// </summary>
        public MonsterDTO() { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public MonsterDTO(IMonster model)
        {
            MonsterId = model.MonsterId;
            MonsterName = model.MonsterName;
            MonsterType = model.MonsterType;
            Hp = model.Hp;
            MaxHp = model.MaxHp;
            Attack = model.Attack;
            DefaultAttack = model.DefaultAttack;
            Speed = model.Speed;
            DefaultSpeed = model.DefaultSpeed;
            Dodge = model.Dodge;
            DefaultDodge = model.DefaultDodge;
            Week = model.Week;
        }
    }
}
