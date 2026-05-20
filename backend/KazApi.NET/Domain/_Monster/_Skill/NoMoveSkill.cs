using KazApi.Common._Log;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// 行動しないスキル
    /// </summary>
    public class NoMoveSkill : ISkill
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public NoMoveSkill(SkillDTO dto)
                    : base(dto) { }

        public override void Use(
            IEnumerable<IMonster> monsters,
            IMonster me,
            ILog<BattleMetaData> logger)
        {
            // 何もしない
        }
    }
}
