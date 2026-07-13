using CSLib.Lib;
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
        public async Task<IActionResult> Init()
        {
            try
            {
                // ドロップダウンの選択肢を取得
                IEnumerable<CodeDTO> dropDown = await _service.FetchDropDown();
                return StatusCode(HttpStatus.OK, dropDown);
            }
            catch (Exception e)
            {
                return StatusCode(HttpStatus.InternalServerError, Message.Create(e));
            }
        }

        [HttpPost("api/edit/fetchMonsters")]
        public async Task<IActionResult> FetctEditMonsters([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                IEnumerable<EditMonsterDTO> monsters = await _service.FetchEditMonsters(loginId);
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
        public async Task<IActionResult> UpdateMonsterStatus([FromBody] IEnumerable<EditMonsterDTO>? monsters)
        {
            if (monsters == null || monsters.Count() == 0) return StatusCode(HttpStatus.BadRequest);

            using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
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
                    await _service.UpdateMonsterStatus(changeMonsters);

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
        public async Task<IActionResult> InitAllMonsterStatus()
        {
            try
            {
                using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
                {
                    await _service.InitAllMonsterStatus();
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
        public async Task<IActionResult> InitAllMonsterSkills()
        {
            try
            {
                using (var transaction = new TransactionScope(TransactionScopeAsyncFlowOption.Enabled))
                {
                    await _service.InitAllMonsterSkills();
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
        public async Task<IActionResult> FecthEditSkills([FromBody] string? loginId)
        {
            if (string.IsNullOrEmpty(loginId)) return StatusCode(HttpStatus.BadRequest);

            try
            {
                IEnumerable<EditSkillsDTO> result = await _service.FecthEditSkills(loginId);
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
        public async Task<IActionResult> FecthAllSkills()
        {
            try
            {
                IEnumerable<AllSkillDTO> result = await _service.FetchAllSkills();
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
        public async Task<IActionResult> UpdateMonsterSkills([FromBody] IEnumerable<EditSkillsDTO>? skills)
        {
            if (skills == null || skills.Count() == 0) return StatusCode(HttpStatus.BadRequest);

            try
            {
                skills = skills.Where(e => e.IsChanged == true);
                await _service.UpdateMonsterSkills(skills);
                
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
