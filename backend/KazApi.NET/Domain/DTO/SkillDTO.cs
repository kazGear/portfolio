using System.Text.Json.Serialization;
using KazApi.Common._Filter;
using KazApi.Domain._Monster._Skill;

namespace KazApi.Domain.DTO
{
    /// <summary>
    /// スキルパラメータクラス
    /// </summary>
    public class SkillDTO
    {
        private string _skillId;
        private string _skillName;
        private int _skillType;
        private int _elementType;
        private int _stateType;
        private int _targetType;
        private int _attack;
        private double _defaultCritical;
        private double _critical;
        private double _hitRate;

        [JsonPropertyName("SkillId")]
        public string SkillId
        {
            get { return _skillId; }
            set { _skillId = Validation.Id(value); }
        }
        
        [JsonPropertyName("SkillName")]
        public string SkillName 
        {
            get { return _skillName; }
            set { _skillName =  Validation.Name(value); }
        }

        
        [JsonPropertyName("SkillType")]
        public int SkillType 
        {
            get { return _skillType; }
            set { _skillType = Validation.SkillType(value); }
        }
        
        [JsonPropertyName("ElementType")]
        public int ElementType 
        {
            get { return _elementType; }
            set { _elementType = Validation.ElementType(value); }
        }
        
        [JsonPropertyName("StateType")]
        public int StateType 
        {
            get { return _stateType; }
            set { _stateType = Validation.StateType(value); }
        }
        
        [JsonPropertyName("TargetType")]
        public int TargetType 
        {
            get { return _targetType; }
            set { _targetType = Validation.TargetType(value); }
        }
        
        [JsonPropertyName("Attack")]
        public int Attack 
        {
            get { return _attack; }
            set { _attack = Validation.Strength(value); }
        }
        
        [JsonPropertyName("Weight")]
        public int Weight { get; set; }

        [JsonPropertyName("DefaultCritical")]
        public double DefaultCritical 
        {
            get { return _defaultCritical; }
            set { _defaultCritical = Validation.Rate(value); } 
        }

        [JsonPropertyName("Critical")]
        public double Critical
        {
            get { return _critical; }
            set { _critical = Validation.Rate(value); }
        }

        [JsonPropertyName("HitRate")]
        public double HitRate 
        {
            get { return _hitRate; }
            set { _hitRate = Validation.Rate(value); }
        }

        [JsonPropertyName("Remarks")]
        public string? Remarks { get; set; } = string.Empty;

        [JsonPropertyName("EffectTime")]
        public int EffectTime { get; set; } = 1000;

        /// <summary>
        /// コンストラクタ
        /// デシリアライズのため必須
        /// </summary>
        public SkillDTO() { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public SkillDTO(ISkill model)
        {
            SkillId = model.SkillId;
            SkillName = model.SkillName;
            SkillType = model.SkillType;
            ElementType = model.ElementType;
            StateType = model.StateType;
            TargetType = model.TargetType;
            Attack = model.Attack;
            Weight = model.Weight;
            DefaultCritical = model.DefaultCritical;
            Critical = model.Critical;
            HitRate = model.HitRate;
            Remarks = "";
        }
    }
}
