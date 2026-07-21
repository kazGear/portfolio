using CSLib.Lib;
using Dapper;
using PrivateApi.Domain.DTO;
using Repository.Repository;
using Repository.Repository.sql;

namespace PrivateApi.Service
{
    public class BatchLogger
    {
        private readonly IDatabase _posgre;

        public BatchLogger(IConfiguration Configuration)
        {
            _posgre = new PostgreSQL(ConnectionString.Get(Configuration));
        }

        public async Task<BatchConfigDTO> InsertStartLog(string batchName)
        {
            var param = new DynamicParameters();
            param.Add("batch_name", batchName);

            // idは自動採番
            await _posgre.Execute(BatchLogSQL.InsertStartLog(), param);

            // 自動採番されたid, configdata取得
            IEnumerable<BatchConfigDTO> configs =
                await _posgre.Select<BatchConfigDTO>(BatchLogSQL.GetConfig(), param);

            return configs.Single();
        }

        public async Task UpdateError(Exception ex, BatchConfigDTO config)
        {
            var param = new DynamicParameters();
            param.Add("status", "ERROR");
            param.Add("message", $"{ex.Message}\n{ex.StackTrace}");
            param.Add("log_id", config.LogId);

            await _posgre.Execute(BatchLogSQL.UpdateStatus(), param);
        }

        public async Task UpdateStatus(BatchConfigDTO config, TimeSpan timeSpan)
        {
            var param = new DynamicParameters();

            param.Add("log_id", config.LogId);
            param.Add("message", "");

            if (config.TimeoutMinutes < timeSpan.TotalMinutes)
            {   
                param.Add("status", "TIMEOUT");
            }
            else if (config.ExpectedDurationMinutes < timeSpan.TotalMinutes)
            {
                param.Add("status", "SLOW");
            }
            else
            {
                param.Add("status", "SUCCESS");
            }
            await _posgre.Execute(BatchLogSQL.UpdateStatus(), param);
        }
    }
}
