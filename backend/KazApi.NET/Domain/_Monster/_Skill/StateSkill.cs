using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain._GameSystem;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// 状態スキルクラス
    /// </summary>
    public class StateSkill : ISkill
    {
        private static readonly IEnumerable<CStateType> POSITIVE_SKILLS = [
            CStateType.AUTOHEAL, CStateType.POWERUP, CStateType.DODGEUP, CStateType.CRITICALUP
            ];
        private IState _state;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public StateSkill(SkillDTO dto, IState state)
                   : base(dto)
        {
            _state = state;
        }

        public override void Use(
            IEnumerable<IMonster> monsters,
            IMonster me,
            ILog<BattleMetaData> logger)
        {
            bool isPositive = POSITIVE_SKILLS.Where(e => e.Value == _state.StateType)
                                             .Count() >= 1;

            if (isPositive) // 有利スキル
            {
                GivePositiveState(me, logger);
            }
            else // 状態異常スキル
            {
                GiveNegativeState(monsters, me, logger);
            }
        }

        /// <summary>
        /// 良い状態を与える
        /// </summary>
        private void GivePositiveState(IMonster me, ILog<BattleMetaData> logger)
        {
            me.AcceptState(_state!.DeepCopy(), this, logger);
        }

        /// <summary>
        /// 悪い状態を与える
        /// </summary>
        private void GiveNegativeState(
            IEnumerable<IMonster> monsters,
            IMonster me,
            ILog<BattleMetaData> logger)
        {
            if (TargetType == CTarget.ENEMY_RANDOM.Value)
            {
                // 単体付与
                IMonster enemy = BattleSystem.SelectOneEnemy(monsters.ToList());

                if (!IsHitSkill(this, enemy)) // 効かないことがある
                {
                    MissLogging(enemy, logger);
                    return;
                }
                enemy.AcceptState(_state.DeepCopy(), this, logger);
            }
            else
            {
                // 全体付与
                foreach (IMonster enemy in monsters)
                {
                    if (!IsHitSkill(this, enemy)) // 効かないことがある
                    {
                        MissLogging(enemy, logger);
                        continue;
                    }
                    enemy.AcceptState(_state.DeepCopy(), this, logger);
                }
            }
        }

        /// <summary>
        /// 状態異常を受けなかった際のログ
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
                $"{enemy.MonsterName}は{this.SkillName}にかからなかった！")
                );
        }
    }
}
