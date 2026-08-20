import CommonFrame from "../components/common/CommonFrame";
import { ChangeEvent, useEffect, useState } from 'react';
import ProjectUsageByFeature from '../components/JobAnalyzePage/ProjectUsageByFeature';
import CommonSelect from '../components/common/CommonSelect';

const JobAnalyzePage = () => {
    const [selectedAnalyze, setSelectedAnalyze] = useState<string>("projectUsageByFeature");

    const selectHandler = (e: ChangeEvent<HTMLSelectElement>) => {
        setSelectedAnalyze(e.currentTarget.value);
    }

    useEffect(() => {
        console.log("useEffect発火", "JobAnalyzePage");
    }, []);

    return (
        <div>
            <CommonFrame styleObj={{margin: "20px", overflowX: "hidden"}}>
                <CommonSelect title="分析内容選択" styleObj={{width: "300px"}} onChange={selectHandler}>
                    <option value="">選択してください</option>
                    <option value="projectUsageByLanguage">採用状況&emsp;（言語別）</option>
                    <option value="projectUsageByFrameworkLibrary">採用状況&emsp;（フレーム、ライブラリ別）</option>
                    <option value="projectUsageByRole">採用状況&emsp;（職能別）</option>
                    <option value="projectUsageByInfrastructure">採用状況&emsp;（インフラ別）</option>
                    <option value="projectUsageByDatabase">採用状況&emsp;（DB別）</option>
                    <option value="projectUsageByCloud">採用状況&emsp;（クラウド別）</option>
                </CommonSelect>
            </CommonFrame>
            <CommonFrame styleObj={{margin: "0px 20px", height: "75vh", overflowX: "hidden"}}>
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

            </CommonFrame>
        </div>
    )
};

export default JobAnalyzePage;
