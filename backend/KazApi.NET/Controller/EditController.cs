using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using KazApi.Domain.DTO;
using System.Transactions;
using KazApi.Domain._Factory;
using KazApi.Service;

namespace KazApi.Controller
{
    [ApiController]
    public class EditController : ControllerBase
    {
        private readonly EditService _service;

        public EditController(IConfiguration configuration)
        {
            _service = new EditService(configuration);
        }

        /// <summary>
        /// 初期処理
        /// </summary>
        [HttpPost("api/edit/init")]
        public ActionResult<string> Init()
        {
            // ドロップダウンの選択肢を取得
            IEnumerable<CodeDTO> dropDown = _service.FetchDropDown();
            return JsonConvert.SerializeObject(dropDown);
        }

        [HttpPost("api/edit/fetchMonsters")]
        public ActionResult<string> FetctEditMonsters([FromQuery] string loginId)
        {
            IEnumerable<EditMonsterDTO> monsters = _service.FetchEditMonsters(loginId);
            return JsonConvert.SerializeObject(monsters);
        }

        /// <summary>
        /// モンスターのステータスを設定する
        /// </summary>
        [HttpPost("api/edit/updateMonsterStatus")]
        public ActionResult UpdateMonsterStatus([FromBody] IEnumerable<EditMonsterDTO> monsters)
        {
            using (TransactionScope transaction = new TransactionScope())
            {
                try
                {
                    IEnumerable<EditMonsterDTO> changeMonsters
                        = monsters.Where(e => e.IsChanged == true);

                    // 未設定値はデフォルト値とする
                    foreach (EditMonsterDTO monster in changeMonsters)
                    {
                        if (monster.AfterName   == null) monster.AfterName = monster.MonsterName;
                        if (monster.AfterHp     == null) monster.AfterHp = monster.Hp;
                        if (monster.AfterAttack == null) monster.AfterAttack = monster.Attack;
                        if (monster.AfterSpeed  == null) monster.AfterSpeed = monster.Speed;
                        if (monster.AfterWeek   == null) monster.AfterWeek = monster.Week;
                    }
                    _service.UpdateMonsterStatus(changeMonsters);

                    transaction.Complete();
                }
                catch (Exception)
                {
                    return StatusCode(500, "モンスターステータスの更新に失敗しました。");
                }
            }
            return Ok(200);
        }

        /// <summary>
        /// 全モンスターのステータスを初期化する
        /// </summary>
        [HttpPost("api/edit/initAllMonsterStatus")]
        public ActionResult InitAllMonsterStatus()
        {
            try
            {
                using (TransactionScope transaction = new TransactionScope())
                {
                    _service.InitAllMonsterStatus();
                    transaction.Complete();
                }
            }
            catch (Exception)
            {
                return StatusCode(500);
            }
            return Ok(200);
        }

        /// <summary>
        /// 全モンスターのスキルを初期化する
        /// </summary>
        [HttpGet("api/edit/initAllMonsterSkills")]
        public ActionResult InitAllMonsterSkills()
        {
            try
            {
                using (TransactionScope transaction = new TransactionScope())
                {
                    _service.InitAllMonsterSkills();
                    transaction.Complete();
                }
            }
            catch (Exception)
            {
                return StatusCode(500);
            }
            return Ok(200);
        }


        /// <summary>
        /// 編集用モンスターデータ（スキル付き）を取得
        /// </summary>
        [HttpPost("api/edit/fecthEditSkills")]
        public ActionResult<string> FecthEditSkills([FromQuery] string loginId)
        {
            try
            {
                IEnumerable<EditSkillsDTO> result = _service.FecthEditSkills(loginId);
                result = new EditFactory().CreateMonstersWithSkills(result);

                return JsonConvert.SerializeObject(result);
            }
            catch (Exception)
            {
                return StatusCode(500);
            }
        }

        /// <summary>
        /// 全スキルを取得
        /// </summary>
        [HttpGet("api/edit/fecthAllSkills")]
        public ActionResult<string> FecthAllSkills()
        {
            IEnumerable<AllSkillDTO> result = _service.FetchAllSkills();
            return JsonConvert.SerializeObject(result);
        }

        /// <summary>
        /// モンスターのスキルを変更する
        /// </summary>
        [HttpPost("api/edit/UpdateMonsterSkills")]
        public ActionResult UpdateMonsterSkills([FromBody] IEnumerable<EditSkillsDTO> skills)
        {
            try
            {
                skills = skills.Where(e => e.IsChanged == true);
                _service.UpdateMonsterSkills(skills);
            }
            catch (Exception)
            {
                return StatusCode(500, "スキルの更新に失敗しました。");
            }
            return Ok(200);
        }
    }
}
