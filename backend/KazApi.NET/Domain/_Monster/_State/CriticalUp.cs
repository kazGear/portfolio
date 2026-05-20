using KazApi.Common._Log;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._State
{
    /// <summary>
    /// クリティカル率アップ状態クラス
    /// </summary>
    public class CriticalUp : IState, IPositiveSkill
    {
        private static readonly double CRITICAL_GAIN = 4.0;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public CriticalUp(StateDTO dto) : base(dto) { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public CriticalUp(string name, string shortName, int stateType, double cancelRate)
                   : base(name, shortName, stateType, cancelRate) { }

        public override IState DeepCopy()
            => new CriticalUp(base.Name, base.ShortName, base.StateType, base.CancelRate);

        public override void DisabledLogging(IMonster me, ILog<BattleMetaData> logger)
        {
            IList<ISkill> result = new List<ISkill>();

            // 全スキルのクリティカル率を戻す
            foreach (ISkill skill in me.CurrentSkills())
            {
                skill.SetCritical(skill.DefaultCritical);
                result.Add(skill);
            }
            me.UpdateSkills(result);

            logger.Logging(new BattleMetaData(
                me.MonsterId,
                base._disabledState,
                base.ShortName,
                $"{me.MonsterName}のクリティカル率が元に戻った。")
                );
        }

        /// <summary>
        /// クリティカル率を上昇させる
        /// </summary>
        public override void Impact(IMonster me, ILog<BattleMetaData> logger)
        {
            IList<ISkill> result = new List<ISkill>();

            // 全スキルのクリティカル率を補正
            foreach (ISkill skill in me.CurrentSkills())
            {
                if (skill.Critical == skill.DefaultCritical)
                {
                    skill.SetCritical(skill.Critical * CRITICAL_GAIN);
                }
                result.Add(skill);
            }
            me.UpdateSkills(result);
        }
    }
}
