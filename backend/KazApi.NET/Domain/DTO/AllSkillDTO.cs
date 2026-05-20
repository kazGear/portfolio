using KazApi.Common._Filter;

namespace KazApi.Domain.DTO
{
    public class AllSkillDTO
    {
        private string _skillId;
        private string _skillName;
        private string _skillTypeName;
        private string _elementName;
        private int _attack;

        public string SkillId
        {
            get {  return _skillId; }
            set { _skillId = Validation.Id(value); }
        }

        public string SkillName
        {
            get { return _skillName; }
            set { _skillName = Validation.Name(value); }
        }

        public string SkillTypeName
        {
            get { return _skillTypeName; }
            set { _skillTypeName = Validation.Name(value); }
        }

        public string ElementName
        {
            get { return _elementName; }
            set { _elementName = Validation.Name(value); }
        }

        public int Attack
        {
            get { return _attack; }
            set { _attack = Validation.Strength(value); }
        }
    }
}
