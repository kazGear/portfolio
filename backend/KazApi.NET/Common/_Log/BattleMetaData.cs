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
        public readonly string TargetMonsterId = string.Empty;
        
        [JsonPropertyName("BeforeHp")]
        public readonly int BeforeHp = NONE;
        
        [JsonPropertyName("ImpactPoint")]
        public readonly int ImpactPoint = NONE;
        
        [JsonPropertyName("StateName")]
        public readonly string StateName = string.Empty;
        
        [JsonPropertyName("EnableState")]
        public readonly bool EnableState = false;
        
        [JsonPropertyName("DisableState")]
        public readonly bool DisableState = false;
        
        [JsonPropertyName("SkillId")]
        public readonly string SkillId = string.Empty;

        [JsonPropertyName("EffectTime")]
        public readonly int EffectTime; // millSecond

        [JsonPropertyName("Message")]
        public readonly string Message = string.Empty;
        
        [JsonPropertyName("IsStop")]
        public readonly bool IsStop = false;

        [JsonPropertyName("IsDodge")]
        public readonly bool IsDodge = false;

        [JsonPropertyName("AllLoser")]
        public readonly bool AllLoser = false;
        
        [JsonPropertyName("ExistWinner")]
        public readonly bool ExistWinner = false;
        
        [JsonPropertyName("WinnerMonsterId")]
        public readonly string WinnerMonsterId = string.Empty;
        
        [JsonPropertyName("WinnerMonsterName")]
        public readonly string WinnerMonsterName = string.Empty;

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
