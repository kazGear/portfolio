using KazApi.Common._Filter;
using KazApi.Domain._Const;
using KazApi.Domain._Monster._State;
using System.Text.Json.Serialization;

namespace KazApi.Domain.DTO
{
    public class StateDTO
    {
        private string _codeId;
        private int _stateType;
        private string _name;
        private string _shortName;
        private double _cancelRate;

        [JsonPropertyName("CodeId")]
        public string CodeId 
        {
            get { return _codeId; }
            set { _codeId = Validation.Id(value); }
        }
      
        [JsonPropertyName("StateType")]
        public int StateType 
        {
            get { return _stateType; }
            set { _stateType = Validation.StateType(value); }
        }
        
        [JsonPropertyName("Name")]
        public string Name 
        {
            get { return _name; }
            set { _name = Validation.Name(value); }
        }

        [JsonPropertyName("ShortName")]
        public string ShortName
        { 
            get { return _shortName; }
            set { _shortName = Validation.ShortName(value);}
        }

        [JsonPropertyName("CancelRate")]
        public double CancelRate 
        {
            get { return _cancelRate; }
            set { _cancelRate = Validation.Rate(value); }
        }

        [JsonPropertyName("Activate")]
        public bool Activate { get; set; }

        /// <summary>
        /// コンストラクタ
        /// デシリアライズのため必須
        /// </summary>
        public StateDTO() { }

        /// <summary>
        /// コンストラクタ
        /// </summary>
        public StateDTO(IState model)
        {
            CodeId = CCodeType.STATE.Value;
            StateType = model.StateType;
            Name = model.Name;
            ShortName = model.ShortName;
            CancelRate = model.CancelRate;
            Activate = model.Activate;
        }
    }
}
