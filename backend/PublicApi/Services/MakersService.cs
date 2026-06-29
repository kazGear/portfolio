using CSLib.Lib;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql;
using Repository.Repository.sql.publicApi;

namespace PublicApi.Services
{
    public class MakersService
    {
        private readonly IDatabase _posgre;

        public MakersService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public IEnumerable<CodeResponse> Get()
        {
            IEnumerable<CodeResponse> makers = _posgre.Select<CodeResponse>(MakersSQL.GetMakers());
            return makers;
        }
    }
}
