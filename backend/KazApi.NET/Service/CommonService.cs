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

        public void UpdateImage(string loginId, string image)
        {
            var param = new
            {
                login_id = loginId,
                image,
            };
            _posgre.Execute(CommonSQL.UpdateUserImage(), param);
        }

        /// <summary>
        /// コード値を取得
        /// </summary>
        public IEnumerable<CodeDTO> FetchElementCode()
            => _posgre.Select<CodeDTO>(CommonSQL.FetchElementCode());

    }
}
