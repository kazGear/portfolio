using KazApi.Common._Filter;
using System.Text.Json.Serialization;

namespace KazApi.Domain.DTO
{
    public class EditSkillsDTO
    {
        // モンスターステータス

        private string _itemId;
        private string _monsterId;
        private string _monsterName;
        private int _hp;
        private int _monsterAttack;
        private int _speed;
        private string _weekName;
        private int _skillAttack;

        // TODO バリデーション組み込み
        [JsonPropertyName("ItemId")]
        public string ItemId
        {
            get { return _itemId; }
            set {  _itemId = Validation.Id(value); }
        }

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

        [JsonPropertyName("Hp")]
        public int Hp
        {
            get { return _hp; }
            set { _hp = Validation.Hp(value); }
        }

        [JsonPropertyName("MonsterAttack")]
        public int MonsterAttack
        {
            get { return _monsterAttack; }
            set { _monsterAttack = Validation.Strength(value); }
        }

        [JsonPropertyName("Speed")]
        public int Speed 
        {
            get { return _speed; }
            set { _speed = Validation.Strength(value); }
        }

        [JsonPropertyName("WeekName")]
        public string WeekName
        {
            get { return _weekName; }
            set { _weekName = Validation.Name(value); }
        }

        [JsonPropertyName("MySkillId")]
        public string? MySkillId { get; set; }

        [JsonPropertyName("SkillId")]
        public string? SkillId { get; set; }

        [JsonPropertyName("SkillName")]
        public string? SkillName { get; set; }

        [JsonPropertyName("SkillAttack")]
        public int SkillAttack
        {
            get { return _skillAttack; }
            set { _skillAttack = Validation.Strength(value); }
        }

        [JsonPropertyName("SkillElementName")]
        public string? SkillElementName { get; set; }

        [JsonPropertyName("IsChanged")]
        public bool IsChanged { get; set; }

        // 各スキル

        [JsonPropertyName("MySkillIds")]
        public IList<string> MySkillIds { get; set; } = [];

        [JsonPropertyName("SkillIds")]
        public IList<string> SkillIds { get; set; } = [];

        [JsonPropertyName("SkillNames")]
        public IList<string> SkillNames { get; set; } = [];       

        [JsonPropertyName("SkillAttacks")]
        public IList<int> SkillAttacks { get; set; } = [];

        [JsonPropertyName("SkillElementNames")]
        public IList<string> SkillElementNames { get; set; } = [];
    }
}
