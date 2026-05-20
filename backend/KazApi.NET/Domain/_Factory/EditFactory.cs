using KazApi.Domain.DTO;

namespace KazApi.Domain._Factory
{
    /// <summary>
    /// スキル生成クラス
    /// </summary>
    public class EditFactory
    {
        
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public EditFactory()
        {

        }

        /// <summary>
        /// 編集用モンスターデータ（スキル付き）を構築
        /// </summary>
        public IEnumerable<EditSkillsDTO> CreateMonstersWithSkills(
            IEnumerable<EditSkillsDTO> monsters)
        {
            IList<EditSkillsDTO> result = [];
            EditSkillsDTO dto = new();

            for (int i = 0; i < monsters.Count(); i++)
            {
                // スキルリスト構築
                dto.MySkillIds.Add(monsters.ElementAt(i).MySkillId!);
                dto.SkillIds.Add(monsters.ElementAt(i).SkillId!);
                dto.SkillNames.Add(monsters.ElementAt(i).SkillName!);
                dto.SkillAttacks.Add(monsters.ElementAt(i).SkillAttack);
                dto.SkillElementNames.Add(monsters.ElementAt(i).SkillElementName!);

                // 最終レコード or モンスターが変わる直前
                if (   i == monsters.Count() - 1
                    || monsters.ElementAt(i).MonsterId != monsters.ElementAt(i + 1).MonsterId)
                {
                    // ステータス構築
                    dto.ItemId = monsters.ElementAt(i).ItemId;
                    dto.MonsterId = monsters.ElementAt(i).MonsterId;
                    dto.MonsterName = monsters.ElementAt(i).MonsterName;
                    dto.Hp = monsters.ElementAt(i).Hp;
                    dto.MonsterAttack = monsters.ElementAt(i).MonsterAttack;
                    dto.Speed = monsters.ElementAt(i).Speed;
                    dto.WeekName = monsters.ElementAt(i).WeekName;
                    
                    result.Add(dto);
                    dto = new();
                }
            }
            return result;
        }
    }
}
