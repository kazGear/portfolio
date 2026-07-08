using PrivateApi.Common._Log;
using PrivateApi.Domain.DTO;

namespace PrivateApi.Domain._Monster._Skill
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
