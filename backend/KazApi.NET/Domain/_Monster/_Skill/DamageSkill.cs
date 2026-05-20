using CSLib.Lib;
using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain._GameSystem;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// ダメージ付与クラス
    /// </summary>
    public class DamageSkill : ISkill
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public DamageSkill(SkillDTO dto)
                    : base(dto) { }

        public override void Use(IEnumerable<IMonster> monsters, IMonster me, ILog<BattleMetaData> logger)
        {
            int target = OneOrAll();

            if (target == CTarget.ENEMY_RANDOM.Value) // 単体攻撃
            {
                IMonster enemy = BattleSystem.SelectOneEnemy(monsters);
                AttackEnemy(enemy, me, logger);
            }
            else if (target == CTarget.ENEMY_ALL.Value) // 全体攻撃
            {
                // 全体攻撃は威力弱め
                PowerDown();
                monsters = monsters.Where(e => e.Hp > 0);
                foreach (IMonster enemy in monsters) AttackEnemy(enemy, me, logger);
                InitPower();
            }
            else
            {
                throw new Exception("不適切なターゲットタイプです。");
            }
        }
        /// <summary>
        /// 敵を攻撃
        /// </summary>
        private void AttackEnemy(IMonster enemy, IMonster me, ILog<BattleMetaData> logger)
        {
            // 攻撃回避
            if (!IsHitSkill(this, enemy))
            {
                MissLogging(enemy, logger);
                return;
            }

            // ダメージ量が多少揺れる
            int damage = new URandom().RandomChangeInt(Attack + me.Attack, CSysRate.PHYSICAL_SKILL_DAMAGE.Value);

            // 弱点等のダメージ欲正
            damage = base.WeeknessDamage(this, enemy, damage, logger);
            damage = base.CriticalDamage(this, damage, logger);

            HitLogging(enemy, damage, logger);
            enemy.AcceptDamage(damage);
        }
        /// <summary>
        /// 攻撃がヒットした際のログ
        /// </summary>
        private void HitLogging(IMonster enemy, int damage, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                enemy.MonsterId,
                enemy.Hp,
                damage,
                SkillId,
                EffectTime,
                $"{enemy.MonsterName}は{damage}のダメージを受けた。")
                );
        }
        /// <summary>
        /// 攻撃を外した際際のログ
        /// </summary>
        private void MissLogging(IMonster enemy, ILog<BattleMetaData> logger)
        {
            int noDamage = 0;
            bool isDodge = true;

            logger.Logging(new BattleMetaData(
                enemy.MonsterId,
                enemy.Hp,
                noDamage,
                base.SkillId,
                base.EffectTime,
                isDodge,
                $"{enemy.MonsterName}は攻撃を回避した！")
                );
        }
        /// <summary>
        /// 単体攻撃か全体攻撃かを選択する
        /// </summary>
        private int OneOrAll()
        {
            if (base.TargetType == CTarget.ENEMY_RANDOM_OR_ALL.Value)
            {
                return new URandom().RandomBool() ? CTarget.ENEMY_RANDOM.Value
                                                  : CTarget.ENEMY_ALL.Value;
            }
            return base.TargetType;
        }
    }
}
