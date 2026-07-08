using PrivateApi.Common._Filter;

namespace PrivateApi.Domain.DTO
{
    public class MonsterTypeDTO
    {
        private string _monsterTypeName;

        public int MonsterTypeId { get; set; }

        public string MonsterTypeName
        { 
            get { return _monsterTypeName; }
            set { _monsterTypeName = Validation.Name(value); } 
        }
    }
}
