import useApiErrorHandler from "../hooks/useApiErrorHandler";
import CommonNowLoading from "../components/common/CommonNowLoading";
import CommonFrame from "../components/common/CommonFrame";

const JobAnalyzePage = () => {

    const errorHandler = useApiErrorHandler();

    return (
        <div>
            <h1 style={{background: "white", paddingLeft: "40px"}}>案件分析ページ</h1>
            <CommonFrame styleObj={{margin: "0px 20px", height: "75vh"}}>
                <h1 style={{padding: "40px"}}>実装中</h1>
            </CommonFrame>
        </div>
    )
};

export default JobAnalyzePage;
