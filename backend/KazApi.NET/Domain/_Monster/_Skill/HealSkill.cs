using CSLib.Const;
using CSLib.Lib;
using PrivateApi.Common._Log;
using PrivateApi.Domain.DTO;

namespace PrivateApi.Domain._Monster._Skill
{
    /// <summary>
    /// 回復スキルクラス
    /// </summary>
    public class HealSkill : ISkill, IPositiveSkill
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public HealSkill(SkillDTO dto)
                  : base(dto) { }

        public override void Use(
            IEnumerable<IMonster> monsters,
            IMonster me,
            ILog<BattleMetaData> logger)
        {
            int healPoint = new Randoms().RandomChangeInt(
                (Attack + me.Attack), CSysRate.MAGIC_SKILL_DAMAGE.Value
                );

            // MaxHp以上に回復はできない
            int healAble = me.MaxHp - me.Hp;
            healPoint = healAble < healPoint ? healAble : healPoint;

            logger.Logging(new BattleMetaData(
                me.MonsterId,
                me.Hp,
                healPoint * -1,
                SkillId,
                EffectTime,
                $"{me.MonsterName}は{healPoint}ポイント回復した！"
                ));

            // マイナス数値の減算でHP加算
            me.AcceptDamage(healPoint * -1);
        }
    }
}
