using CSLib.Lib;
using Dapper;
using PublicApi.RequestDtos;
using PublicApi.ResponseDtos;
using Repository.Repository;
using Repository.Repository.sql.publicApi;
using System.Text;

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
            // SQL パーツ構築
            string conditions = CreateConditions(req);
            string order      = CreateOrder(req);
            string sort       = CreateSortTarget(req, order);

            DynamicParameters param = CreateParams(req);

            // ギター情報取得
            IReadOnlyList<GuitarResponse> guitars =
                _posgre.Select<GuitarResponse>(GuitarsSQL.GetGuitars(conditions, sort), param).ToList();
            
            // 検索総件数
            int totalCount = _posgre.Select<int>(GuitarsSQL.GetTotalCount(conditions), param).First();

            GuitarsResponse res = new GuitarsResponse()
            {
                TotalCount = totalCount,
                Page       = req.Page,
                PageSize   = req.PageSize,
                TotalPages = (int)Math.Ceiling((decimal)totalCount / (decimal)req.PageSize),
                HasPrev    = req.Page > 1,
                HasNext    = req.Page * req.PageSize < totalCount,
                Guitars    = guitars,
            };
            return res;
        }

        private DynamicParameters CreateParams(GuitarsRequest req)
        {
            var param = new DynamicParameters();

            param.Add("maker", req.MakerCd);
            param.Add("name", string.IsNullOrWhiteSpace(req.Name) ? null : $"%{req.Name}%");
            param.Add("series", string.IsNullOrWhiteSpace(req.Series) ? null : $"%{req.Series}%");
            param.Add("color_cd", req.ColorCd);
            param.Add("body_material_top_cd", req.BodyMaterialTopCd);
            param.Add("body_material_back_cd", req.BodyMaterialBackCd);
            param.Add("min_price", req.MinPrice);
            param.Add("max_price", req.MaxPrice);
            param.Add("page", req.Page);
            param.Add("page_size", req.PageSize);

            return param;
        }

        private string CreateConditions(GuitarsRequest req)
        {
            StringBuilder conditions = new StringBuilder();

            if (req.MakerCd != null) 
            {
                conditions.AppendLine("AND maker = @maker");
            }
            if (!string.IsNullOrWhiteSpace(req.Name))
            {
                conditions.AppendLine($"AND name ilike @name");
            }
            if (!string.IsNullOrWhiteSpace(req.Series))
            {
                conditions.AppendLine("AND series ilike @series");
            }
            if (req.ColorCd != null)
            {
                conditions.AppendLine("AND color_cd = @color_cd");
            }
            if (req.BodyMaterialTopCd != null && req.BodyMaterialTopCd >= 0)
            {
                conditions.AppendLine("AND body_material_top = @body_material_top_cd");
            }
            if (req.BodyMaterialBackCd != null && req.BodyMaterialBackCd >= 0)
            {
                conditions.AppendLine("AND body_material_back = @body_material_back_cd");
            }
            if (req.MinPrice != null)
            {
                conditions.AppendLine("AND price >= @min_price");
            }
            if (req.MaxPrice != null)
            {
                conditions.AppendLine("AND price <= @max_price");
            }
            return conditions.ToString();
        }

        private string CreateOrder(GuitarsRequest req)
        {
            // 基本は昇順
            return req.Order == "DESC" ? "DESC" : "ASC";
        }

        private string CreateSortTarget(GuitarsRequest req, string order)
        {
            string sortResult = string.Empty;

            if (req.Sort == "maker")
            {
                sortResult = $" ORDER BY maker {order}, name ASC ";
            }
            else if (req.Sort == "name")
            {
                sortResult = $" ORDER BY name {order} ";
            }
            else if (req.Sort == "price")
            {
                sortResult = $" ORDER BY price {order}, name ASC  ";
            }
            else
            {
                sortResult = $" ORDER BY maker ASC, name ASC ";
            }
            return sortResult;
        }
    }
}
