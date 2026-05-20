namespace KazApi.Domain._Monster._Skill
{
    /// <summary>
    /// マーカーインターフェイス
    /// 有利な効果を及ぼすスキル
    /// 
    /// Use(IEnumerable<IMonster> monsters, IMonster me)として
    /// 敵モンスターが渡ってくるが、ポジティブスキルにおいて敵は行動しない。
    /// </summary>
    public interface IPositiveSkill
    {

    }
}
