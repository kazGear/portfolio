using CSLib.Lib;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql.publicApi;

namespace PublicApi.Services
{
    public class ColorsService
    {
        private readonly IDatabase _posgre;

        public ColorsService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public IEnumerable<CodeResponse> Get()
        {
            IEnumerable<CodeResponse> colors = _posgre.Select<CodeResponse>(ColorsSQL.GetColors());
            return colors;
        }
    }
}
