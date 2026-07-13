using CSLib.Lib;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class CommonService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public CommonService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(configuration));
        }

        public async Task UpdateImage(string loginId, string image)
        {
            var param = new
            {
                login_id = loginId,
                image,
            };
            await _posgre.Execute(CommonSQL.UpdateUserImage(), param);
        }

        /// <summary>
        /// コード値を取得
        /// </summary>
        public async Task<IEnumerable<CodeDTO>> FetchElementCode()
            => await _posgre.Select<CodeDTO>(CommonSQL.FetchElementCode());

    }
}
