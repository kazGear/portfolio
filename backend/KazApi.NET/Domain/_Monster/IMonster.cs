 using KazApi.Common._Log;
using KazApi.Domain._Const;
using KazApi.Domain._GameSystem;
using KazApi.Domain._Monster._Skill;
using KazApi.Domain._Monster._State;
using KazApi.Domain.DTO;

namespace KazApi.Domain._Monster
{
    /// <summary>
    /// モンスターインターフェイス
    /// </summary>
    public abstract class IMonster
    {
        protected ISet<IState> _status = new HashSet<IState>();
        protected IList<ISkill> _skills = new List<ISkill>();

        public string MonsterId { get; protected set; }
        public string MonsterName { get; protected set; }
        public int MonsterType { get; protected set; }
        public int Hp { get; protected set; }
        public int MaxHp { get; protected set; } = 0;
        public int Attack { get; protected set; }
        public int DefaultAttack { get; protected set; }
        public int Speed { get; protected set; }
        public int DefaultSpeed { get; protected set; }
        public double Dodge { get; protected set; }
        public double DefaultDodge { get; protected set; }
        public int Team { get; protected set; }
        public int Week { get; protected set; }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public IMonster(MonsterDTO dto, IEnumerable<ISkill> skills, IEnumerable<IState> status)
        {
            MonsterId = dto.MonsterId;
            MonsterName = dto.MonsterName;
            MonsterType = dto.MonsterType;
            Hp = dto.Hp;
            if (MaxHp == 0) MaxHp = dto.MaxHp;
            Attack = dto.Attack;
            DefaultAttack = dto.DefaultAttack;
            Speed = dto.Speed;
            DefaultSpeed= dto.Speed;
            Dodge = dto.Dodge;
            DefaultDodge = dto.DefaultDodge;

            foreach (IState state in status) _status.Add(state);
            foreach (ISkill skill in skills) _skills.Add(skill);
            Team = CTeam.UNKNOWN.Value;
            Week = dto.Week;
        }
        /// <summary>
        /// 行動する
        /// </summarys>
        public void Move(IList<IMonster> monsters, ILog<BattleMetaData> logger)
        {
            // 戦闘不能中は攻撃不可
            if (Hp <= 0) return;
            // 攻撃対象外
            if (monsters.Where(e => e.Hp > 0).Count() <= 0) return;

            ISkill skill = SelectSkill();

            if (skill is IPositiveSkill)
            {
                PositiveMove(monsters, skill, logger);
            }
            else
            {
                AttackMove(monsters, skill, logger);
            }
        }
        /// <summary>
        /// 攻撃する
        /// </summary>
        protected abstract void AttackMove(
            IEnumerable<IMonster> monsters,
            ISkill skill,
            ILog<BattleMetaData> logger
            );

        /// <summary>
        /// 有利な行動
        /// </summary>
        protected abstract void PositiveMove(
            IEnumerable<IMonster> monsters,
            ISkill skill,
            ILog<BattleMetaData> logger
            );

        /// <summary>
        /// スキルを選択
        /// </summary>
        public abstract ISkill SelectSkill();
        /// <summary>
        /// 現在のスキルを見る
        /// </summary>
        public IEnumerable<ISkill> CurrentSkills() => new List<ISkill>(_skills);
        /// <summary>
        /// 現在のステータスを見る
        /// </summary>
        public IEnumerable<IState> CurrentStatus() => new List<IState>(_status);
        /// <summary>
        /// ステータスを更新する
        /// </summary>
        public void UpdateStatus(ISet<IState> changedStatus) => _status = changedStatus;
        /// <summary>
        /// スキルを更新する
        /// </summary>
        public void UpdateSkills(IList<ISkill> changedSkills) => _skills = changedSkills;
        /// <summary>
        /// ダメージを受ける
        /// </summary>
        public void AcceptDamage(int damage) => Hp -= (Hp < damage) ? Hp : damage;
        /// <summary>
        /// チームを決定する
        /// </summary>
        public void DefineTeam(int team) => Team = team;
        /// <summary>
        /// 攻撃力を変更する
        /// </summary>
        public void SetAttack(int attack) => Attack = attack;
        /// <summary>
        /// スピードを変更する
        /// </summary>
        public void SetSpeed(int speed) => Speed = speed;
        /// <summary>
        /// 回避力を変更する
        /// </summary>
        public void SetDodge(double dodge) => Dodge = dodge;
        /// <summary>
        /// 攻撃力を元に戻す
        /// </summary>
        public void InitAttack() => Attack = DefaultAttack;
        /// <summary>
        /// スピードを元に戻す
        /// </summary>
        public void InitSpeed() => Speed = DefaultSpeed;
        /// <summary>
        /// 回避力を元に戻す
        /// </summary>
        public void InitDodge() => Dodge = DefaultDodge;
        /// <summary>
        /// 状態異常になる
        /// </summary>
        public void AcceptState(IState state, ISkill skill, ILog<BattleMetaData> logger)
        {
            // 状態異常は重複しない
            if (_status.Where(e => e.GetStateType() == state.GetStateType()).Count() <= 0)
            {
                bool enableState = true;

                _status.Add(state);
                logger.Logging(new BattleMetaData(
                    MonsterId,
                    skill.SkillId,
                    skill.EffectTime,
                    state.ShortName,
                    enableState,
                    $"{MonsterName}は{state.Name}状態になった。")
                    );
            }
            else
            {
                logger.Logging(new BattleMetaData(MonsterId, $"{MonsterName}は既に{state.Name}状態になっている。"));
            }
        }
        /// <summary>
        /// 状態異常解除・未解除の振り分け
        /// </summary>
        public void RefreshStatus(ILog<BattleMetaData> logger)
        {
            IEnumerable<IState> currentStatus = this.CurrentStatus();
            ISet<IState> changedStatus = new HashSet<IState>();

            foreach (IState state in currentStatus)
            {
                if (BattleSystem.StateIsDisabled(state) && state.Activate)
                {
                    // ステータス解除
                    state.DisabledLogging(this, logger);
                }
                else
                {
                    // ステータス有効化・継続
                    state.Activation();
                    changedStatus.Add(state);
                }
            }
            UpdateStatus(changedStatus);
        }
        /// <summary>
        /// 状態異常の効果を受ける
        /// </summary>
        public void StateImpact(ILog<BattleMetaData> logger)
        {
            // 変化したコレクションの操作は例外となる > 新しいコレクションを操作して回避
            IEnumerable<IState> copyStatus = new List<IState>(_status);
            foreach (IState state in copyStatus) state.Impact(this, logger);
        }
        /// <summary>
        /// 行動できる状態か判定
        /// </summary>
        public bool IsMoveAble()
        {
            int notMove = _status.Where(e => e is IDisableMove)
                                 .Count();
            return notMove <= 0;
        }
    }
}