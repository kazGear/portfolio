using KazApi.Domain._Const;
using System.Text.RegularExpressions;

namespace KazApi.Common._Filter
{
    public static class Validation
    {
        private static string GetMessageUndefined(string target, object actual)
        {
            return $"定義されていない{target}です。>>> {actual}";
        }
        private static string GetMessageSetWithinRange(string target, int min, int max)
        {
            return $"{target}は{min}-{max}の範囲で設定してください。";
        }
        private static string GetMessageSetWithinRange(string target, double min, double max)
        {
            return $"{target}は{min}-{max}の範囲で設定してください。";
        }
        private static string GetMessageNotMatching(string target, object actual)
        {
            return $"{target}のパターンが異なっています。>>> {actual}";
        }

        public static int Amount(int amount)
        {
            if (amount < 0)
                throw new Exception($"金額は{amount}円以上で設定してください。");

            return amount;
        }

        public static int BattleReportSerial(int serial)
        {
            int min = 1;
            int max = 6;

            string message = GetMessageSetWithinRange("バトルレポートの連番", min, max);

            if (serial < min) throw new Exception(message);
            if (max < serial) throw new Exception(message);

            return serial;
        }

        public static int CodeValue(int codeValue)
        {
            int min = 0;
            int max = 99;

            string message = GetMessageSetWithinRange("コード値", min, max);

            if (codeValue < min) throw new Exception(message);
            if (max < codeValue) throw new Exception(message);

            return codeValue;
        }

        public static int Count(int count)
        {
            if (count < 0)
                throw new Exception($"カウントは{count}以上で設定してください。");

            return count;
        }

        public static int ElementType(int week)
        {
            if (week == -1) return week; // tmp

            IReadOnlyCollection<int> values = CElement.GetValues();

            if (!values.Contains(week))
                throw new Exception(GetMessageUndefined("属性", week));

            return week;
        }

        public static double Rate(double rate)
        {
            double min = 0.0;
            double max = 1.0;

            string message = GetMessageSetWithinRange("割合", min, max);

            if (rate < min) throw new Exception(message);
            if (max < rate) throw new Exception(message);

            return rate;
        }

        public static int Hp(int hp)
        {
            int min = 0;
            int max = 999;

            string message = GetMessageSetWithinRange("HP", min, max);

            if (hp < min) throw new Exception(message);
            if (max < hp) throw new Exception(message);

            return hp;
        }

        public static string Id(string itemId)
        {
            string pattern = @"^[a-zA-Z]+\d{3}$";

            if (!Regex.IsMatch(itemId, pattern))
                throw new Exception(GetMessageNotMatching("ID", itemId));

            return itemId;
        }

        public static string LoginId(string loginId)
        {
            string pattern = @"^[a-zA-Z0-9-_]{4,15}$";

            if (!Regex.IsMatch(loginId, pattern))
                throw new Exception(GetMessageNotMatching("ログインID", loginId));

            return loginId;
        }

        public static string LoginPass(string loginPass)
        {
            int minLength = 4;

            if (minLength > loginPass.Length)
                throw new Exception($"ログインパスワードは{minLength}文字以上で設定してください。");

            return loginPass;

        }

        public static string MonsterType(string monsterType)
        {
            IReadOnlyCollection<string> values = CMonsterType.GetValues();

            if (!values.Contains(monsterType))
                throw new Exception(GetMessageUndefined("モンスタータイプ", monsterType));

            return monsterType;
        }

        public static string MySkillId(string mySkillId)
        {
            string pattern = @"^myskill\d{4}$";

            if (!Regex.IsMatch(mySkillId, pattern))
                throw new Exception(GetMessageNotMatching("マイスキルID", mySkillId));

            return mySkillId;
        }

        public static string Name(string name)
        {
            int maxLength = 15;

            if (maxLength < name.Length)
                throw new Exception($"名称は{maxLength}文字以内で設定してください。");

            return name;
        }

        public static string ShortName(string shortName)
        {
            int maxLength = 5;

            if (maxLength < shortName.Length)
                throw new Exception($"省略名は{maxLength}文字以内で設定してください。");

            return shortName;
        }

        public static int SkillType(int skillType)
        {
            // tmp
            if (skillType == -1) return skillType;

            IReadOnlyCollection<int> values = CSkillType.GetValues();

            if (!values.Contains(skillType))
                throw new Exception(GetMessageUndefined("スキル", skillType));

            return skillType;
        }

        public static int StateType(int stateType)
        {
            if (stateType == -1) return stateType; // tmp

            IReadOnlyCollection<int> values = CStateType.GetValues();

            if (!values.Contains(stateType))
                throw new Exception(GetMessageUndefined("状態", stateType));

            return stateType;
        }

        public static int Strength(int speed)
        {
            int min = 0;
            int max = 255;

            string message = GetMessageSetWithinRange("素早さ", min, max);

            if (speed < min) throw new Exception(message);
            if (max < speed) throw new Exception(message);

            return speed;
        }

        public static int TargetType(int targetType)
        {
            if (targetType == -1) return targetType; // tmp

            IReadOnlyCollection<int> values = CTarget.GetValues();

            if (!values.Contains(targetType))
                throw new Exception(GetMessageUndefined("ターゲット", targetType));

            return targetType;
        }
    }
}