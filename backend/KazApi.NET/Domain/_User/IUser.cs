namespace KazApi.Domain._User;

public abstract class IUser
{
    public string UserName { get; protected set; }
    public string Password { get; protected set; }
    public int Money { get; protected set; }
    public int Hp { get; protected set; }
    public int Mp { get; protected set; }
    public int Attack { get; protected set; }
    public int Defense { get; protected set; }
    public bool IsInvalid { get; protected set; }
}
