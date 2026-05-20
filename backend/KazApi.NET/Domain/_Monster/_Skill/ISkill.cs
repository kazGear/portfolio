using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// スキルインターフェイス
    /// </summary>
    public abstract class ISkill
    {
        protected int _initialAttack;

        public string SkillId { get; protected set; }
        public string SkillName { get; protected set; }
        public int SkillType { get; protected set; }
        public int Attack { get; protected set; }
        public int ElementType { get; protected set; }
        public int StateType { get; protected set; }
        public int TargetType { get; protected set; }
        public int Weight { get; protected set; }
        public double DefaultCritical { get; protected set; }
        public double Critical { get; protected set; }
        public double HitRate { get; protected set; }
        public int EffectTime { get; protected set; }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public ISkill(SkillDTO dto)
        {
            SkillId = dto.SkillId;
            SkillName = dto.SkillName;
            SkillType = dto.SkillType;
            Attack = dto.Attack;
            _initialAttack = dto.Attack;
            ElementType = dto.ElementType;
            StateType = dto.StateType;
            TargetType = dto.TargetType;
            Weight = dto.Weight;
            DefaultCritical = dto.DefaultCritical;
            Critical = dto.Critical;
            HitRate = dto.HitRate;
            EffectTime = dto.EffectTime;
        }

        /// <summary>
        /// スキルを使用
        /// </summary>
        public abstract void Use(IEnumerable<IMonster> monsters, IMonster me, ILog<BattleMetaData> logger);

        /// <summary>
        /// 全体攻撃の際は威力減少
        /// </summary>
        public void PowerDown()
        {
            double allAttackDamage = Attack * new URandom().RandomDouble(
                CSysRate.ALL_ATTACK_ADJUST_MIN.Value,
                CSysRate.ALL_ATTACK_ADJUST_MAX.Value
                );
            Attack = (int)allAttackDamage;
        }

        /// <summary>
        /// 攻撃力を初期値に戻す
        /// </summary>
        protected void InitPower()
            => Attack = _initialAttack;

        /// <summary>
        /// クリティカル率を変更する
        /// </summary>
        /// <param name="newRate"></param>
        public void SetCritical(double newRate)
            => Critical = newRate;

        /// <summary>
        /// 弱点属性によるダメージの算出
        /// </summary>
        public int WeeknessDamage(
            ISkill skill,
            IMonster enemy,
            int damage,
            ILog<BattleMetaData> logger)
        {
            // 弱点ダメージが発生しようがなければダメージの変動なし
            if (skill.ElementType == CElement.NONE.Value) return damage;
            if (enemy.Week == CElement.NONE.Value) return damage;

            if (enemy.Week == skill.ElementType)
            {
                damage = (int)(damage * CSysRate.WEEK_DAMAGE.Value);
                logger.Logging(new BattleMetaData("弱点ダメージ！"));
            }
            return damage;
        }

        /// <summary>
        /// クリティカルによるダメージ
        /// </summary>
        public int CriticalDamage(ISkill skill, int damage, ILog<BattleMetaData> logger)
        {
            double randomVal = new URandom().RandomDouble(0.0, 1.0);
            bool isCritical = randomVal <= skill.Critical;

            if (isCritical)
            {
                damage = (int)(damage * CSysRate.CRITICAL_DAMAGE.Value);
                logger.Logging(new BattleMetaData("クリティカルヒット！"));
            }
            return damage;
        }

        /// <summary>
        /// スキルが命中したか判定
        /// true: hit, false: miss
        /// </summary>
        public bool IsHitSkill(ISkill skill, IMonster enemy)
        {
            bool result = false;
            double randVal = new URandom().RandomDouble(0.0, 1.0);

            if (randVal <= (skill.HitRate - enemy.Dodge)) result = true;

            return result;
        }

        /// <summary>
        /// DTOへ変換
        /// </summary>
        public static IEnumerable<SkillDTO> ModelToDTO(IEnumerable<ISkill> models)
        {
            IList<SkillDTO> result = [];

            foreach (ISkill model in models)
                result.Add(new SkillDTO(model));

            return result;
        }
    }
}
