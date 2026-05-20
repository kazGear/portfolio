using KazApi.Domain.DTO;
using KazApi.Repository;
using KazApi.Repository.sql;

namespace KazApi.Service
{
    public class CommonService
    {
        private readonly IDatabase _posgre;

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public CommonService(IConfiguration configuration)
        {
            _posgre = new PostgreSQL(configuration);
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
