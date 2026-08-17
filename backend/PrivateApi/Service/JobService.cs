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

        public async Task<JobsResponse> GetJobs(JobsRequest req)
        {
            // SQL パーツ構築
            string conditions        = CreateConditions(req);
            string featureConditions = CreateFeatureConditions(req);

            DynamicParameters param = CreateParams(req);

            // 案件情報取得
            IEnumerable<JobDTO> jobs =
                await _posgre.Select<JobDTO>(
                    JobSQL.SelectJobs(conditions, featureConditions), param);

            // 検索総件数
            int totalCount =
                (await _posgre.Select<int>(
                    JobSQL.GetTotalCount(conditions, featureConditions), param)).First();

            // csvを分離する
            SplitFeatures(jobs);
            SplitOptions(jobs);

            JobsResponse res = new JobsResponse()
            {
                TotalCount = totalCount,
                Page       = req.Page,
                PageSize   = req.PageSize,
                TotalPages = (int)Math.Ceiling((decimal)totalCount / (decimal)req.PageSize),
                HasPrev    = req.Page > 1,
                HasNext    = req.Page * req.PageSize < totalCount,
                Jobs       = jobs,
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

            param.Add("title", $"%{req.Title}%");
            param.Add("location", req.Location);
            param.Add("min_salary_at_month_specified_min", req.MinSalaryAtMonthSpecifiedMin);
            param.Add("min_salary_at_month_specified_max", req.MinSalaryAtMonthSpecifiedMax);
            param.Add("max_salary_at_month_specified_min", req.MaxSalaryAtMonthSpecifiedMin);
            param.Add("max_salary_at_month_specified_max", req.MaxSalaryAtMonthSpecifiedMax);
            param.Add("Work_place", req.WorkPlace);
            param.Add("source_site", req.SourceSite);
            param.Add("page", req.Page);
            param.Add("page_size", req.PageSize);
            param.Add("is_hide_old_job", req.IsHideOldJob);

            foreach (string feature in req.FeatureNames)
            {
                param.Add(feature, feature);
            }
            return param;
        }

        private string CreateConditions(JobsRequest req)
        {
            StringBuilder conditions = new StringBuilder();

            if (!string.IsNullOrWhiteSpace(req.Title))
            {
                conditions.AppendLine("AND title iLIKE @title");
            }
            if (!string.IsNullOrWhiteSpace(req.Location))
            {
                conditions.AppendLine($"AND location = @location");
            }
            if (!string.IsNullOrWhiteSpace(req.MinSalaryAtMonthSpecifiedMin))
            {
                conditions.AppendLine("AND min_salary_at_month >= @min_salary_at_month_specified_min::int");
            }
            if (!string.IsNullOrWhiteSpace(req.MinSalaryAtMonthSpecifiedMax))
            {
                conditions.AppendLine("AND min_salary_at_month <= @min_salary_at_month_specified_max::int");
            }
            if (!string.IsNullOrWhiteSpace(req.MaxSalaryAtMonthSpecifiedMin))
            {
                conditions.AppendLine("AND max_salary_at_month >= @max_salary_at_month_specified_min::int");
            }
            if (!string.IsNullOrWhiteSpace(req.MaxSalaryAtMonthSpecifiedMax))
            {
                conditions.AppendLine("AND max_salary_at_month <= @max_salary_at_month_specified_max::int");
            }
            if (!string.IsNullOrWhiteSpace(req.WorkPlace))
            {
                conditions.AppendLine($"AND work_place = @work_place");
            }
            if (!string.IsNullOrWhiteSpace(req.SourceSite))
            {
                conditions.AppendLine($"AND source_site = @source_site");
            }
            if (req.IsHideOldJob)
            {
                // 掲載するのは更新日（取得日）から○○ヵ月以内の案件
                conditions.AppendLine($"AND NOW() - updated_at <= '3 month'");
            }
            return conditions.ToString();
        }

        private string CreateFeatureConditions(JobsRequest req)
        {
            StringBuilder conditions = new StringBuilder();

            foreach (string feature in req.FeatureNames)
            {
                string SQL = @$"
                    AND EXISTS
                        (
                           SELECT 1
                             FROM t_job_features AS f
                            WHERE v.id = f.job_id
                              AND f.feature_name = @{feature}
                        )
                ";
                conditions.AppendLine(SQL);
            }

            return conditions.ToString();
        }

        public async Task<IEnumerable<string>> GetFeatures(string category)
        {
            var param = new DynamicParameters();
            param.Add("category", category);

            IEnumerable<string> features =
                await _posgre.Select<string>(JobSQL.GetFeatures(), param);
            
            return features;
        }
    }
}
