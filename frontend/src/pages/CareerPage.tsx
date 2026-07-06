import meImage from "../../../_docs/images/me.png";
import {
    career,
    littleSkills,
    metaData,
    portfolio,
    prPoint,
    profile,
    readBooks,
    skills,
    specialty,
    useTools
} from "../data/Career";
import { calcAge } from "../lib/CommonLogic";
import Name from "../components/careerPage/Name";
import Age from "../components/careerPage/Age";
import Address from "../components/careerPage/Address";
import Skills from "../components/careerPage/Skills";
import LowSkills from "../components/careerPage/LowSkills";
import Tools from "../components/careerPage/Tools";
import Specialty from "../components/careerPage/Specialty";
import ReadBooks from "../components/careerPage/ReadBooks";
import PrPoint from "../components/careerPage/PrPoint";
import Portfolio from "../components/careerPage/Portfolio";
import Career from "../components/careerPage/Career";
import styled from "styled-components";
import { useState } from "react";
import Button from "../components/common/Button";
import { COLORS } from "../lib/Constants";

const SleftSide = styled.div`
    width: 40%;
    height: 100vh;
    overflow-y: scroll;
    background: rgb(58, 58, 58);
    color: aliceblue;
`;

const SrightSide = styled.div`
    width: 60%;
    height: 100vh;
    overflow-y: scroll;
    background: aliceblue;
    color: rgb(58, 58, 58);
`;

const SinnerFrame = styled.div`
    margin: 20px;
`;

const SdetailFilter = styled.div`
    position: absolute;
    z-index: 1000;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(58, 58, 58, 0.75);
`;

const SdetailModal = styled.div`
    position: absolute;
    z-index: 2000;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 80%;
    height: 80%;
    border-radius: 15px;
    color: ${COLORS.MAIN_FONT_COLOR};
    background: aliceblue;
    overflow-y: scroll;
`;

const SbuttonFrame = styled.div`
    position: absolute;
    right: 20px;
    bottom: 20px;
`;


const CareerPage = () => {
    const [isShowDetail, setIsShowDetail] = useState(false)

    return (
        <main style={{
            color: "skyblue",
            fontFamily: "BIZ UD明朝 Medium, Consolas",
            display: "flex"
            }}>
            <SleftSide>
                <SinnerFrame>
                    <div style={{textAlign: "center"}}>
                        <h1>職務経歴書</h1>
                        <img src={meImage}
                             alt="me"
                             style={{width: "200px", borderRadius: "15px"}}/>
                        <p>最終更新日：{metaData.lastUpdate}</p>
                    </div>

                    <div>
                        <Button text="詳細情報を見る"
                                onClick={() => setIsShowDetail(true)}
                                styleObj={{width: "150px", height: "40px"}}/>
                    </div>

                    {/* 基本情報 */}
                    <div>
                        <h3 >プロフィール</h3>
                        <table>
                            <tbody>
                                <Name name={profile.myName}/>
                                <Age age={calcAge("1987/08/23")}/>
                                <Address address={profile.myAddress}/>
                                <Skills skills={skills}/>
                                <LowSkills lowSkills={littleSkills}/>
                                <Tools tools={useTools}/>
                            </tbody>
                        </table>
                    </div>

                    {/* ポートフォリオ */}
                    <div >
                        <Portfolio portfolio={portfolio}/>
                    </div>

                    <div>
                        <Button text="詳細情報を見る"
                                onClick={() => setIsShowDetail(true)}
                                styleObj={{width: "150px", height: "40px"}}/>
                    </div>
                </SinnerFrame>
            </SleftSide>

            {/* 経歴部 */}
            <SrightSide>
                <SinnerFrame>
                    <div >
                        <Career career={career}/>
                    </div>
                </SinnerFrame>
            </SrightSide>

            {/* PR部 モーダル*/}
            {
                isShowDetail ?
                (
                    <SdetailFilter>
                        <SdetailModal>
                            <SinnerFrame>
                                <div>
                                    <Specialty specialty={specialty}/>
                                    <br/>
                                    <PrPoint prPoint={prPoint}/>
                                    <br/>
                                    <ReadBooks readBooks={readBooks}/>
                                </div>

                            </SinnerFrame>
                        </SdetailModal>
                        <SbuttonFrame>
                            <Button text="閉じる"
                                    onClick={() => setIsShowDetail(false)}
                                    styleObj={{width: "150px", height: "40px"}}/>
                        </SbuttonFrame>
                    </SdetailFilter>
                )
                : "" // 表示なし
            }
        </main>
    );
};
export default CareerPage;