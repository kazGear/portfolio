using KazApi.Domain._Const;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Factory
{
    /// <summary>
    /// スキル生成クラス
    /// </summary>
    public class StateFactory
    {
        private IEnumerable<CodeDTO> _codeEntities;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public StateFactory(IEnumerable<CodeDTO> codeEntities)
        {
            // 状態コード取得
            _codeEntities = codeEntities;
        }

        /// <summary>
        /// 状態オブジェクトを作成（状態コードより）
        /// </summary>
        public IState Create(int stateCode)
        {
            CodeDTO param = _codeEntities.Where(e => e.Value == stateCode).Single();

            if (stateCode == CStateType.NONE.Value)
                return new None(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.POISON.Value)
                return new Poison(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.DEADLY_POISON.Value)
                return new DeadlyPoison(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.SLEEP.Value)
                return new Sleep(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.CHARM.Value)
                return new Charm(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.SLOW.Value)
                return new Slow(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.POWERUP.Value)
                return new PowerUp(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.DODGEUP.Value)
                return new DodgeUp(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.CRITICALUP.Value)
                return new CriticalUp(param.Name, param.ShortName, param.Value, param.Param3);

            else if (stateCode == CStateType.AUTOHEAL.Value)
                return new AutoHeal(param.Name, param.ShortName, param.Value, param.Param3);

            else
                throw new Exception("存在しない状態コードです。");
        }
        /// <summary>
        /// 状態オブジェクトを作成（DTOより）
        /// </summary>
        public IState Create(int stateCode, StateDTO dto)
        {
            CodeDTO param = _codeEntities.Where(e => e.Value == stateCode).Single();

            if (stateCode == CStateType.NONE.Value)
                return new None(dto);

            else if (stateCode == CStateType.POISON.Value)
                return new Poison(dto);

            else if (stateCode == CStateType.DEADLY_POISON.Value)
                return new DeadlyPoison(dto);

            else if (stateCode == CStateType.SLEEP.Value)
                return new Sleep(dto);

            else if (stateCode == CStateType.CHARM.Value)
                return new Charm(dto);

            else if (stateCode == CStateType.SLOW.Value)
                return new Slow(dto);

            else if (stateCode == CStateType.POWERUP.Value)
                return new PowerUp(dto);

            else if (stateCode == CStateType.DODGEUP.Value)
                return new DodgeUp(dto);

            else if (stateCode == CStateType.CRITICALUP.Value)
                return new CriticalUp(dto);

            else if (stateCode == CStateType.AUTOHEAL.Value)
                return new AutoHeal(dto);

            else
                throw new Exception("存在しない状態コードです。");
        }
    }
}
