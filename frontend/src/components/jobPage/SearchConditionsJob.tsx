import { GuitarParams, GuitarsResponse } from "../../types/Guitar";
import { Code } from "../../types/Code";
// import SearchMaker from "./SearchMaker";
// import SearchColor from "./SearchColor";
// import SearchSeries from "./SearchSeries";
// import SearchName from "./SearchName";
// import SearchMaterialTop from "./SearchMaterialTop";
// import SearchMaterialBack from "./SearchMaterialBack";
// import SearchMinPrice from "./SearchMinPrice";
// import SearchMaxPrice from "./SearchMaxPrice";
import styled from "styled-components";
// import SelectorOrder from "./SelectorOrder";
// import SelectorSort from "./SelectorSort";
// import SelectorPage from "./SelectorPage";
// import SelectorPageSize from "./SelectorPageSize";
import CommonBorderTr from "../common/CommonBorderTr";
import { JobParams, JobsResponse } from "../../types/Job";
import SearchTitle from "./SearchTitle";
import SearchLocation from "./SearchLocation";
import SearchWorkPlace from "./SearchWorkPlace";
import SearchMinSalarySpecifiedMax from "./SearchMinSalarySpecifiedMax";
import SearchMinSalarySpecifiedMin from "./SearchMinSalarySpecifiedMin";
import SearchMaxSalarySpecifiedMin from "./SearchMaxSalarySpecifiedMin";
import SearchMaxSalarySpecifiedMax from "./SearchMaxSalarySpecifiedMax";
import SearchSourceSite from "./SearchSourceSite";
import SearchLanguage from "./SearchLanguage";
import SelectorPageSize from "./SelectorPageSize";
import SelectorPage from "./SelectorPage";

const Th = styled.th`
    text-align: left;
    min-width: 80px;
    font-size: 14px;
    font-weight: bolder;
`;
const Td = styled.td`
    text-align: left;
    font-size: 14px;
    font-weight: bolder;
`;

const styleObj = {
    margin: "5px 20px",
}

interface ArgProps {
    jobsRes:        JobsResponse | null;
    jobParams:      JobParams;
    languages:      string[] | null;
    searchHandler: (jParams: JobParams) => Promise<void>;
}

const SearchConditionsJob = ({jobsRes,
                              jobParams,
                              languages,
                              searchHandler
}: ArgProps) => {

    return (
        <div style={{margin: "10px", overflow: "hidden"}}>
            <table>
                <thead>
                    <CommonBorderTr>
                        <Th>検索方法</Th>
                        <td style={{fontSize: "14px", paddingLeft: "20px"}}>
                            ※自動検索<br/>検索条件を変更すると<br/>自動的に検索されます。
                        </td>
                    </CommonBorderTr>
                </thead>
                <tbody>
                    <CommonBorderTr>
                        <Th>タイトル</Th>
                        <Td>
                            <SearchTitle jobParams={jobParams}
                                         searchHandler={searchHandler}
                                         styleObj={{margin: "7px 0px 7px 20px"}} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>所在地</Th>
                        <Td><SearchLocation jobParams={jobParams} searchHandler={searchHandler} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>勤務地</Th>
                        <Td><SearchWorkPlace jobParams={jobParams} searchHandler={searchHandler} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最低報酬</Th>
                        <Td>
                            <SearchMinSalarySpecifiedMin jobParams={jobParams}
                                                         searchHandler={searchHandler}
                                                         styleObj={styleObj}/>
                            <br/>&emsp;&emsp;～<br/>
                            <SearchMinSalarySpecifiedMax jobParams={jobParams}
                                                         searchHandler={searchHandler}
                                                         styleObj={styleObj}/>
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最高報酬</Th>
                        <Td>
                            <SearchMaxSalarySpecifiedMin jobParams={jobParams}
                                                         searchHandler={searchHandler}
                                                         styleObj={styleObj}/>
                            <br/>&emsp;&emsp;～<br/>
                            <SearchMaxSalarySpecifiedMax jobParams={jobParams}
                                                         searchHandler={searchHandler}
                                                         styleObj={styleObj}/>
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ソースサイト</Th>
                        <Td><SearchSourceSite jobParams={jobParams} searchHandler={searchHandler} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ページサイズ</Th>
                        <Td>
                            <SelectorPageSize jobParams={jobParams}
                                              searchHandler={searchHandler}
                                              styleObj={styleObj}/>
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>選択ページ</Th>
                        <Td>
                            <SelectorPage jobParams={jobParams}
                                          jobsRes={jobsRes}
                                          searchHandler={searchHandler}
                                          styleObj={{margin: "0px 0px 6px 15px"}}/>
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>言語</Th>
                        <Td>
                            <SearchLanguage jobParams={jobParams}
                                            languages={languages}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                </tbody>
            </table>
        </div>
    );
}

export default SearchConditionsJob;