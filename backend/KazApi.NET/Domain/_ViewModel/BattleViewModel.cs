using PrivateApi.Common._Log;
using PrivateApi.Domain.DTO;

namespace PrivateApi.Domain._ViewModel
{
    public class BattleViewModel
    {
        public IEnumerable<MonsterDTO> Monsters { get; set; }
        public IEnumerable<BattleMetaData> BattleLog { get; set; }
    }
}
