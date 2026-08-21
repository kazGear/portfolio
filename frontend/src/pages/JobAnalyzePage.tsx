import CommonFrame from "../components/common/CommonFrame";
import { ChangeEvent, useState } from 'react';
import ProjectUsageByFeature from '../components/JobAnalyzePage/ProjectUsageByFeature';
import CommonSelect from '../components/common/CommonSelect';
import WorkPlaceByPrefecture from "../components/JobAnalyzePage/WorkPlaceByPrefecture";
import SalaryRangeByFeature from "../components/JobAnalyzePage/SalaryRangeByFeature";
import SavedJobDataStatus from "../components/JobAnalyzePage/SavedJobDataStatus";

const JobAnalyzePage = () => {
    const [selectedAnalyze, setSelectedAnalyze] = useState<string>("projectUsageByFeature");

    const selectHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        setSelectedAnalyze(e.currentTarget.value);
    }

    return (
        <div>
            <CommonFrame styleObj={{margin: "20px", overflowX: "hidden"}}>
                <CommonSelect title="分析内容選択" styleObj={{width: "300px"}} onChange={selectHandler}>
                    <option value="">選択してください</option>

                    <option value="workPlaceByPrefecture">勤務形態&emsp;（都道府県別）</option>

                    <option value="savedJobDataStatus">案件保存状態&emsp;（運用者用）</option>

                    <option value="projectUsageByLanguage">採用状況&emsp;（言語別）</option>
                    <option value="projectUsageByFrameworkLibrary">採用状況&emsp;（フレーム＆ライブラリ別）</option>
                    <option value="projectUsageByRole">採用状況&emsp;（職能別）</option>
                    <option value="projectUsageByInfrastructure">採用状況&emsp;（インフラ別）</option>
                    <option value="projectUsageByDatabase">採用状況&emsp;（DB別）</option>
                    <option value="projectUsageByCloud">採用状況&emsp;（クラウド別）</option>

                    <option value="salaryRangeByLanguage">給与レンジ&emsp;（言語別）</option>
                    <option value="salaryRangeByFrameworkLibrary">給与レンジ&emsp;（フレーム＆ライブラリ別）</option>
                    <option value="salaryRangeByRole">給与レンジ&emsp;（職能別）</option>
                    <option value="salaryRangeByInfrastructure">給与レンジ&emsp;（インフラ別）</option>
                    <option value="salaryRangeByDatabase">給与レンジ&emsp;（DB別）</option>
                    <option value="salaryRangeByCloud">給与レンジ&emsp;（クラウド別）</option>
                </CommonSelect>
            </CommonFrame>
            <CommonFrame styleObj={{margin: "0px 20px", height: "75vh", overflowX: "hidden"}}>
                { selectedAnalyze === "workPlaceByPrefecture" ? <WorkPlaceByPrefecture /> : "" }


                { selectedAnalyze === "savedJobDataStatus" ? <SavedJobDataStatus /> : "" }


                { selectedAnalyze === "projectUsageByLanguage" ?
                    <ProjectUsageByFeature category="LANGUAGE" /> : "" }

                { selectedAnalyze === "projectUsageByFrameworkLibrary" ?
                    <ProjectUsageByFeature category="FRAMEWORK_LIBRARY" /> : "" }

                { selectedAnalyze === "projectUsageByRole" ?
                    <ProjectUsageByFeature category="ROLE" /> : "" }

                { selectedAnalyze === "projectUsageByInfrastructure" ?
                    <ProjectUsageByFeature category="INFRASTRUCTURE" /> : "" }

                { selectedAnalyze === "projectUsageByDatabase" ?
                    <ProjectUsageByFeature category="DATABASE" /> : "" }

                { selectedAnalyze === "projectUsageByCloud" ?
                    <ProjectUsageByFeature category="CLOUD" /> : "" }


                { selectedAnalyze === "salaryRangeByLanguage" ?
                    <SalaryRangeByFeature category="LANGUAGE" /> : "" }

                { selectedAnalyze === "salaryRangeByFrameworkLibrary" ?
                    <SalaryRangeByFeature category="FRAMEWORK_LIBRARY" /> : "" }

                { selectedAnalyze === "salaryRangeByRole" ?
                    <SalaryRangeByFeature category="ROLE" /> : "" }

                { selectedAnalyze === "salaryRangeByInfrastructure" ?
                    <SalaryRangeByFeature category="INFRASTRUCTURE" /> : "" }

                { selectedAnalyze === "salaryRangeByDatabase" ?
                    <SalaryRangeByFeature category="DATABASE" /> : "" }

                { selectedAnalyze === "salaryRangeByCloud" ?
                    <SalaryRangeByFeature category="CLOUD" /> : "" }
            </CommonFrame>
        </div>
    )
};

export default JobAnalyzePage;
