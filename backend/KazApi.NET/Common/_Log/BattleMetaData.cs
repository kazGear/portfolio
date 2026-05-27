using KazApi.Domain._Monster;
using System.Text.Json.Serialization;

namespace KazApi.Common._Log
{
    /// <summary>
    /// 戦闘ログクラス
    /// </summary>
    public class BattleMetaData
    {
        private static readonly int NONE = 0;

        [JsonPropertyName("TargetMonsterId")]
        public string TargetMonsterId { get; private set; } = string.Empty;
        
        [JsonPropertyName("BeforeHp")]
        public int BeforeHp { get; private set; } = NONE;
        
        [JsonPropertyName("ImpactPoint")]
        public int ImpactPoint { get; private set; } = NONE;
        
        [JsonPropertyName("StateName")]
        public string StateName { get; private set; } = string.Empty;
        
        [JsonPropertyName("EnableState")]
        public bool EnableState { get; private set; } = false;
        
        [JsonPropertyName("DisableState")]
        public bool DisableState { get; private set; } = false;
        
        [JsonPropertyName("SkillId")]
        public string SkillId { get; private set; } = string.Empty;

        [JsonPropertyName("EffectTime")]
        public int EffectTime { get; private set; } // millSecond

        [JsonPropertyName("Message")]
        public string Message { get; private set; } = string.Empty;
        
        [JsonPropertyName("IsStop")]
        public bool IsStop { get; private set; } = false;

        [JsonPropertyName("IsDodge")]
        public bool IsDodge { get; private set; } = false;

        [JsonPropertyName("AllLoser")]
        public bool AllLoser { get; private set; } = false;
        
        [JsonPropertyName("ExistWinner")]
        public bool ExistWinner { get; private set; } = false;
        
        [JsonPropertyName("WinnerMonsterId")]
        public string WinnerMonsterId { get; private set; } = string.Empty;
        
        [JsonPropertyName("WinnerMonsterName")]
        public string WinnerMonsterName { get; private set; } = string.Empty;

        /// <summary>
        /// コンストラクタ1（ログ区切りのマーカー）
        /// </summary>
        public BattleMetaData()
        {
            IsStop = true;
        }

        /// <summary>
        /// コンストラクタ2
        /// </summary>
        public BattleMetaData(string message)
        {
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ3
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ4
        /// </summary>
        public BattleMetaData(
            bool existWinner,
            bool allLoser,
            IMonster? alive
            )
        {
            ExistWinner = existWinner;
            AllLoser = allLoser;

            if (alive != null)
            {
                TargetMonsterId = alive.MonsterId;
                WinnerMonsterId = alive.MonsterId;
                WinnerMonsterName = alive.MonsterName;
            }
            IsStop = true;
        }

        /// <summary>
        /// コンストラクタ5
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            string skillId,
            int effectTime,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            SkillId = skillId;
            EffectTime = effectTime;
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ6
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            bool disableState,
            string stateShortName,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            DisableState = disableState;
            StateName = stateShortName;
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ7
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            string skillId,
            int beforeHp,
            int impactPoint,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            SkillId = skillId;
            BeforeHp = beforeHp;
            ImpactPoint = impactPoint;
            Message = message;
            EffectTime = 1000;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ8
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            string skillId,
            int effectTime,
            string stateName,
            bool enableState,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            SkillId = skillId;
            EffectTime = effectTime;
            StateName = stateName;
            EnableState = enableState;
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ9
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            int beforeHp,
            int impactPoint,
            string skillId,
            int effectTime,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            BeforeHp = beforeHp;
            ImpactPoint = impactPoint;
            SkillId = skillId;
            EffectTime = effectTime;
            Message = message;
            IsStop = false;
        }

        /// <summary>
        /// コンストラクタ10
        /// </summary>
        public BattleMetaData(
            string targetMonsterId,
            int beforeHp,
            int impactPoint,
            string skillId,
            int effectTime,
            bool isDodge,
            string message
            )
        {
            TargetMonsterId = targetMonsterId;
            BeforeHp = beforeHp;
            ImpactPoint = impactPoint;
            SkillId = skillId;
            EffectTime = effectTime;
            IsDodge = isDodge;
            Message = message;
            IsStop = false;
        }
        public override string ToString() => Message;
    }
}
