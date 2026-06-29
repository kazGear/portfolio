using CSLib.Lib;
using PublicApi.RequestDtos;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql.publicApi;

namespace PublicApi.Services
{
    public class GuitarsService
    {
        private readonly IDatabase _posgre;

        public GuitarsService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public GuitarsResponse Get(GuitarsRequest req)
        {
            IEnumerable<CodeResponse> bodyMaterials =
                _posgre.Select<CodeResponse>(GuitarsSQL.GetGuitars());

            string[] conditions = [];




            return null;
        }
    }
}
