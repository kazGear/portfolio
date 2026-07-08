using CSLib.Lib;
using PrivateApi.Common;
using PrivateApi.Domain._Factory;
using PrivateApi.Domain.DTO;
using PrivateApi.Service;
using Microsoft.AspNetCore.Mvc;
using System.Transactions;

namespace PrivateApi.Controller
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
        [HttpGet("api/edit/init")]
        public IActionResult Init()
        {
            try
            {
                // ドロップダウンの選択肢を取得
                IEnumerable<CodeDTO> dropDown = _service.FetchDropDown();
                return StatusCode(HttpStatus.OK, dropDown);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        [HttpPost("api/edit/fetchMonsters")]
        public IActionResult FetctEditMonsters([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                IEnumerable<EditMonsterDTO> monsters = _service.FetchEditMonsters(loginId);
                return StatusCode(HttpStatus.OK, monsters);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// モンスターのステータスを設定する
        /// </summary>
        [HttpPut("api/edit/updateMonsterStatus")]
        public IActionResult UpdateMonsterStatus([FromBody] IEnumerable<EditMonsterDTO>? monsters)
        {
            if (monsters == null || monsters.Count() == 0) return StatusCode(HttpStatus.BadRequest);

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
                    return StatusCode(HttpStatus.OK);
                }
                catch (Exception e)
                {
                    string message = "モンスターステータスの更新に失敗しました。";
                    return StatusCode(HttpStatus.InternalServerError, Message.Create(e, message));
                }
            }
        }

        /// <summary>
        /// 全モンスターのステータスを初期化する
        /// </summary>
        [HttpPut("api/edit/initAllMonsterStatus")]
        public IActionResult InitAllMonsterStatus()
        {
            try
            {
                using (TransactionScope transaction = new TransactionScope())
                {
                    _service.InitAllMonsterStatus();
                    transaction.Complete();
                }
                return StatusCode(HttpStatus.OK);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 全モンスターのスキルを初期化する
        /// </summary>
        [HttpPut("api/edit/initAllMonsterSkills")]
        public IActionResult InitAllMonsterSkills()
        {
            try
            {
                using (TransactionScope transaction = new TransactionScope())
                {
                    _service.InitAllMonsterSkills();
                    transaction.Complete();
                }
                return StatusCode(HttpStatus.OK);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 編集用モンスターデータ（スキル付き）を取得
        /// </summary>
        [HttpPost("api/edit/fecthEditSkills")]
        public IActionResult FecthEditSkills([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                IEnumerable<EditSkillsDTO> result = _service.FecthEditSkills(loginId);
                result = new EditFactory().CreateMonstersWithSkills(result);

                return StatusCode(HttpStatus.OK, result);

            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// 全スキルを取得
        /// </summary>
        [HttpGet("api/edit/fecthAllSkills")]
        public IActionResult FecthAllSkills()
        {
            try
            {
                IEnumerable<AllSkillDTO> result = _service.FetchAllSkills();
                return StatusCode(HttpStatus.OK, result);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        /// <summary>
        /// モンスターのスキルを変更する
        /// </summary>
        [HttpPut("api/edit/UpdateMonsterSkills")]
        public IActionResult UpdateMonsterSkills([FromBody] IEnumerable<EditSkillsDTO>? skills)
        {
            if (skills == null || skills.Count() == 0) return StatusCode(HttpStatus.BadRequest);

            try
            {
                skills = skills.Where(e => e.IsChanged == true);
                _service.UpdateMonsterSkills(skills);
                
                return StatusCode(HttpStatus.OK);
            }
            catch (Exception e)
            {
                string message = "スキルの更新に失敗しました。";
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e, message));
            }
        }
    }
}
