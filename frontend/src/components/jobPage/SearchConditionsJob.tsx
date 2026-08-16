import styled from "styled-components";
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
import SearchFeatures from "./SearchFeatures";
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
    jobsRes:          JobsResponse | null;
    jobParams:        JobParams;
    languages:        string[] | null;
    frameworkLibrary: string[] | null;
    role:             string[] | null;
    infrastructure:   string[] | null;
    database:         string[] | null;
    cloud:            string[] | null;
    searchHandler: (jParams: JobParams) => Promise<void>;
}

const SearchConditionsJob = ({jobsRes,
                              jobParams,
                              languages,
                              frameworkLibrary,
                              role,
                              infrastructure,
                              database,
                              cloud,
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
                            <SearchFeatures jobParams={jobParams}
                                            features={languages}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ロール</Th>
                        <Td>
                            <SearchFeatures jobParams={jobParams}
                                            features={role}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>インフラ</Th>
                        <Td>
                            <SearchFeatures jobParams={jobParams}
                                            features={infrastructure}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>DB</Th>
                        <Td>
                            <SearchFeatures jobParams={jobParams}
                                            features={database}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>クラウド</Th>
                        <Td>
                            <SearchFeatures jobParams={jobParams}
                                            features={cloud}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>frame and library</Th>
                        <Td>
                            <SearchFeatures jobParams={jobParams}
                                            features={frameworkLibrary}
                                            searchHandler={searchHandler} />
                        </Td>
                    </CommonBorderTr>
                </tbody>
            </table>
        </div>
    );
}

export default SearchConditionsJob;