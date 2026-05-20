using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// 魅了状態クラス
    /// </summary>
    public class Charm : IState, IDisableMove
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Charm(StateDTO dto) : base(dto) 
        {
            base.Activate = true;
        }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public Charm(string name, string shortName, int stateType, double cancelRate)
              : base(name, shortName, stateType, cancelRate)
        {
            base.Activate = true;
        }

        public override IState DeepCopy()
            => new Charm(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster monster, ILog<BattleMetaData> logger)
        {
            logger.Logging(new BattleMetaData(
                monster.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{monster.MonsterName}は我に返った！")
                );
        }

        /// <summary>
        /// 自身に攻撃
        //  有利な効果は使用しない
        /// </summary>
        public override void Impact(IMonster me, ILog<BattleMetaData> logger)
        {
            // 睡眠時は発動しない
            int sleepCnt = me.CurrentStatus()
                             .Where(e => e.StateType == CStateType.SLEEP.Value)
                             .Count();

            if (sleepCnt >= 1) return;

            ISkill skill = me.SelectSkill();
            while (skill is HealSkill || skill is NoMoveSkill) // 選び直し
                skill = me.SelectSkill();

            // 自傷
            logger.Logging(new BattleMetaData(me.MonsterId, $"{me.MonsterName}は自分に攻撃！"));
            logger.Logging(new BattleMetaData(me.MonsterId, $"{me.MonsterName}は {skill.SkillName} を放った！"));
            skill.Use([me], me, logger);
        }
    }
}
