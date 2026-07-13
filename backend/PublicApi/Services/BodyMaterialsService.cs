using CSLib.Lib;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql.publicApi;

namespace PublicApi.Services
{
    public class BodyMaterialsService
    {
        private readonly IDatabase _posgre;

        public BodyMaterialsService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public async Task<IEnumerable<CodeResponse>> Get()
        {
            IEnumerable<CodeResponse> bodyMaterials =
                await _posgre.Select<CodeResponse>(BodyMaterialsSQL.GetBodyMaterials());
            return bodyMaterials;
        }
    }
}
