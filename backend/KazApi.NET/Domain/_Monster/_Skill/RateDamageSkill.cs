using KazApi.Common._Log;
using KazApi.Domain._GameSystem;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// 割合ダメージクラス
    /// </summary>
    public class RateDamageSkill : ISkill
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public RateDamageSkill(SkillDTO dto)
                        : base(dto) { }

        public override void Use(
            IEnumerable<IMonster> monsters,
            IMonster me,
            ILog<BattleMetaData> logger)
        {
            IMonster enemy = BattleSystem.SelectOneEnemy(monsters);

            // 現HPの割合ダメージ
            double damage = enemy.Hp * (Attack / 100.0);

            logger.Logging(new BattleMetaData(
                enemy.MonsterId,
                enemy.Hp,
                (int)damage,
                SkillId,
                EffectTime,
                $"{enemy.MonsterName}は{(int)damage}のダメージを受けた。")
                );

            enemy.AcceptDamage((int)damage);
        }
    }
}
