using CSLib.Lib;
using PublicApi.RequestDtos;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql.publicApi;

namespace PublicApi.Services
{
    public class SeriesService
    {
        private readonly IDatabase _posgre;

        public SeriesService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public async Task<IEnumerable<CodeResponse>> Get(SeriesRequest req)
        {
            object param = new { maker = req.MakerCd };

            IEnumerable<CodeResponse> series =
                await _posgre.Select<CodeResponse>(SeriesSQL.GetSeries(), param);

            return series;
        }
    }
}
