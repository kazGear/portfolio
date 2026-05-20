using KazApi.Domain._Monster;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Factory
{
    /// <summary>
    /// モンスター生成クラス
    /// </summary>
    public class MonsterFactory
    {
        /// <summary>
        /// コンストラクタ
        /// </summary>
        public MonsterFactory()
        {

        }
 
        /// <summary>
        /// モンスターとスキルのマッピング（DTO）
        /// </summary>
        public IEnumerable<MonsterDTO> MappingToMonsterDTO(
            IEnumerable<MonsterDTO> monstersDTO,
            IEnumerable<SkillDTO> skillsDTO,
            IEnumerable<MonsterSkillDTO> monsterSkillsDTO
            )
        {
            IList<MonsterDTO> result = [];

            foreach (MonsterDTO monster in monstersDTO)
            {
                // モンスターのスキル対応表を取得
                IEnumerable<MonsterSkillDTO> skillMap =
                    monsterSkillsDTO.Where(e => e.MonsterId == monster.MonsterId);

                IList<SkillDTO> bindSkills = [];

                foreach (MonsterSkillDTO ms in skillMap)
                {
                    // 対応表を元にスキルを設定
                    SkillDTO skill = skillsDTO.Where(e => e.SkillId == ms.SkillId).Single();
                    bindSkills.Add(skill);
                }
                // スキルを持ったモンスター
                monster.Skills = bindSkills;
                monster.Status = []; // 初期値は状態なし
                result.Add(monster);
            }
            return result;
        }

        /// <summary>
        /// 戦闘用モンスターオブジェクトを構築
        /// </summary>
        public IEnumerable<IMonster> CreateBattleMonsters(
            IEnumerable<MonsterDTO> monsters,
            IEnumerable<CodeDTO> codes)
        {
            SkillFactory skillFactory = new SkillFactory(codes);
            StateFactory stateFactory = new StateFactory(codes);

            // バトルモンスター構築
            IList<IMonster> battleMonsters = [];
            foreach (MonsterDTO m in monsters)
            {
                IEnumerable<ISkill> skills = skillFactory.Create(m.Skills);

                ISet<IState> status = new HashSet<IState>();
                foreach (StateDTO state in m.Status)
                {
                    // 同じ状態は追加しない
                    status.Add(stateFactory.Create(state.StateType, state));
                }
                // スキル、ステータスを設定
                IMonster battleMonster = new Monster(m, skills, status);
                battleMonsters.Add(battleMonster);
            }
            return battleMonsters;
        }

        /// <summary>
        /// モデルからDTOへ変換
        /// </summary>
        public IEnumerable<MonsterDTO> ConvertToDTO(IEnumerable<IMonster> battleMonsters)
        {
            IList<MonsterDTO> monstersDTO = [];
            foreach (IMonster m in battleMonsters)
            {
                IEnumerable<SkillDTO> skillsDTO = ISkill.ModelToDTO(m.CurrentSkills());
                IEnumerable<StateDTO> statusDTO = IState.ModelToDTO(m.CurrentStatus());

                MonsterDTO monsterDTO = new MonsterDTO(m);
                monsterDTO.Skills = skillsDTO;
                monsterDTO.Status = statusDTO;

                monstersDTO.Add(monsterDTO);
            }
            return monstersDTO;
        }
    }
}
