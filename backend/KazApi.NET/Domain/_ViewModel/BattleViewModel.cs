using KazApi.Common._Log;
using KazApi.Domain.DTO;

namespace KazApi.Domain._ViewModel
{
    public class BattleViewModel
    {
        public IEnumerable<MonsterDTO> Monsters { get; set; }
        public IEnumerable<BattleMetaData> BattleLog { get; set; }
    }
}
