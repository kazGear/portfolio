using CSLib.Lib;
using Dapper;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;
using System.Text;

namespace PrivateApi.Service
{
    public class JobService
    {
        private readonly IDatabase _posgre;

        public JobService(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public async Task<JobsResponse> Get(JobsRequest req)
        {
            // SQL パーツ構築
            string conditions = CreateConditions(req);
            //string order = CreateOrder(req);
            //string sort = CreateSortTarget(req, order);

            DynamicParameters param = CreateParams(req);

            // 案件情報取得
            IEnumerable<JobDTO> jobs =
                await _posgre.Select<JobDTO>(JobSQL.SelectJobs()/*, param*/);

            // 検索総件数
            int totalCount =
                (await _posgre.Select<int>(JobSQL.GetTotalCount(conditions), param)).First();

            // csvを分離する
            SplitFeatures(jobs);
            SplitOptions(jobs);

            JobsResponse res = new JobsResponse()
            {
                TotalCount = totalCount,
                Page = req.Page,
                PageSize = req.PageSize,
                TotalPages = (int)Math.Ceiling((decimal)totalCount / (decimal)req.PageSize),
                HasPrev = req.Page > 1,
                HasNext = req.Page * req.PageSize < totalCount,
                Jobs = jobs,
            };
            return res;
        }

        private void SplitFeatures(IEnumerable<JobDTO> jobs)
        {
            foreach (JobDTO job in jobs)
            {
                IEnumerable<string> featuresNames = job.FeatureNamesCSV.Split(",");
                job.FeatureNames = featuresNames;
            }
        }

        private void SplitOptions(IEnumerable<JobDTO> jobs)
        {
            foreach (JobDTO job in jobs)
            {
                IEnumerable<string> options = job.OptionsCSV.Split(",");
                job.Options = options;
            }
        }

        private DynamicParameters CreateParams(JobsRequest req)
        {
            var param = new DynamicParameters();

            //param.Add("maker", req.MakerCd);
            //param.Add("name", string.IsNullOrWhiteSpace(req.Name) ? null : $"%{req.Name}%");
            //param.Add("series", string.IsNullOrWhiteSpace(req.Series) ? null : $"%{req.Series}%");
            //param.Add("color_cd", req.ColorCd);
            //param.Add("body_material_top_cd", req.BodyMaterialTopCd);
            //param.Add("body_material_back_cd", req.BodyMaterialBackCd);
            //param.Add("min_price", req.MinPrice);
            //param.Add("max_price", req.MaxPrice);
            //param.Add("page", req.Page);
            //param.Add("page_size", req.PageSize);

            return param;
        }

        private string CreateConditions(JobsRequest req)
        {
            StringBuilder conditions = new StringBuilder();

            //if (req.MakerCd != null)
            //{
            //    conditions.AppendLine("AND maker = @maker");
            //}
            //if (!string.IsNullOrWhiteSpace(req.Name))
            //{
            //    conditions.AppendLine($"AND guitars.name ilike @name");
            //}
            //if (!string.IsNullOrWhiteSpace(req.Series))
            //{
            //    conditions.AppendLine("AND series ilike @series");
            //}
            //if (req.ColorCd != null)
            //{
            //    conditions.AppendLine("AND color_cd = @color_cd");
            //}
            //if (req.BodyMaterialTopCd != null && req.BodyMaterialTopCd >= 0)
            //{
            //    conditions.AppendLine("AND body_material_top = @body_material_top_cd");
            //}
            //if (req.BodyMaterialBackCd != null && req.BodyMaterialBackCd >= 0)
            //{
            //    conditions.AppendLine("AND body_material_back = @body_material_back_cd");
            //}
            //if (req.MinPrice != null)
            //{
            //    conditions.AppendLine("AND price >= @min_price");
            //}
            //if (req.MaxPrice != null)
            //{
            //    conditions.AppendLine("AND price <= @max_price");
            //}
            return conditions.ToString();
        }

        //private string CreateOrder(GuitarsRequest req)
        //{
        //    // 基本は昇順
        //    return req.Order == "DESC" ? "DESC" : "ASC";
        //}

        //private string CreateSortTarget(GuitarsRequest req, string order)
        //{
        //    string sortResult = string.Empty;

        //    if (req.Sort == "maker")
        //    {
        //        sortResult = $" ORDER BY maker {order}, name ASC ";
        //    }
        //    else if (req.Sort == "name")
        //    {
        //        sortResult = $" ORDER BY name {order} ";
        //    }
        //    else if (req.Sort == "price")
        //    {
        //        sortResult = $" ORDER BY price {order}, name ASC  ";
        //    }
        //    else
        //    {
        //        sortResult = $" ORDER BY name ASC ";
        //    }
        //    return sortResult;
        //}
    }
}
