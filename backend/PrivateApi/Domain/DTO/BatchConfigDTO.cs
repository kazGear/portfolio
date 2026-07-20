using PrivateApi.Common._Filter;

namespace PrivateApi.Domain.DTO
{
    public class BatchConfigDTO
    {
        private long _logId;
        private DateTime _startAt;
        private string _batchName;
        private TimeOnly _scheduledTime;
        private int _startDelayMinutes;
        private int _expectedDurationMinutes;
        private int _timeoutMinutes;

        public long LogId
        {
            get {  return _logId; }
            set { _logId = value; }
        }

        public DateTime StartAt
        {
            get { return _startAt; }
            set { _startAt = value; }
        }

        public string BatchName
        {
            get { return _batchName; }
            set { _batchName = Validation.Name(value); }
        }

        public TimeOnly ScheduledTime
        {
            get { return _scheduledTime; }
            set { _scheduledTime = value; }
        }

        public int StartDelayMinutes
        {
            get { return _startDelayMinutes; }
            set { _startDelayMinutes = Validation.Minute(value); }
        }

        public int ExpectedDurationMinutes
        {
            get { return _expectedDurationMinutes; }
            set { _expectedDurationMinutes = Validation.Minute(value); }
        }

        public int TimeoutMinutes
        {
            get { return _timeoutMinutes; }
            set { _timeoutMinutes = Validation.Minute(value); }
        }
    }
}
