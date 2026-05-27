namespace KazApi.Common
{
    public record ReqMonsterReport(int monsterTypeId, int sortType, bool isAscOrder);

    public record ReqBattleReport(int battleScale, string? from, string? to);

    public record ReqUserRegist(string loginId, string password, string dispName, string dispShortName);

    public record ReqMyItem(string loginId, string itemId);

    public record ReqLogin(string loginId, string password);

    public record ReqShopItem(string loginId, string? selectedShop);

    public record ReqBattleInit(string selectMonstersCount, string loginId);

    public record ReqUserResults(string  betMonsterId,
                                 int     betGil,
                                 decimal betRate,
                                 string  winningMonsterId,
                                 string  loginId);
}
