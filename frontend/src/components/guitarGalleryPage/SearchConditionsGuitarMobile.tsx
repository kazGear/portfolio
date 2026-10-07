import { GuitarParams } from "../../types/Guitar";
import { Code } from "../../types/Code";
import SearchMaker from "./SearchMaker";
import SearchColor from "./SearchColor";
import SearchSeries from "./SearchSeries";
import SearchName from "./SearchName";
import SearchMinPrice from "./SearchMinPrice";
import SearchMaxPrice from "./SearchMaxPrice";
import styled from "styled-components";
import SelectorOrder from "./SelectorOrder";
import SelectorSort from "./SelectorSort";
import CommonBorderTr from "../common/CommonBorderTr";

const Th = styled.th`
    text-align: left;
    min-width: 80px;
    font-size: 13px;
    font-weight: bolder;
`;
const Td = styled.td`
    text-align: left;
    font-size: 13px;
    font-weight: bolder;
`;

const styleObj = {
    margin: "5px 20px",
}

interface ArgProps {
    guitarParams: GuitarParams;
    makers:       Code[] | null;
    series:       Code[] | null;
    colors:       Code[] | null;
}

const SearchConditionsGuitarMobile = ({
    guitarParams,
    makers,
    series,
    colors}: ArgProps
) => {
    const gParams = guitarParams;

    return (
        <div style={{margin: "10px", overflowX: "hidden"}}>
            <table style={{margin: "20px 0px"}}>
                <thead>
                    <CommonBorderTr styleObj={{borderTop: "none"}}>
                        <th style={{textAlign: "left"}}>検索条件</th>
                        <td style={{paddingLeft: "20px", fontWeight: "bold"}}>設定値</td>
                    </CommonBorderTr>
                </thead>
                <tbody>
                    <CommonBorderTr>
                        <Th>メーカー</Th>
                        <Td><SearchMaker guitarParams={gParams} makers={makers} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>カラー</Th>
                        <Td><SearchColor guitarParams={gParams} colors={colors} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>シリーズ</Th>
                        <Td><SearchSeries guitarParams={gParams} series={series} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ギター名</Th>
                        <Td><SearchName guitarParams={gParams} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最低価格</Th>
                        <Td><SearchMinPrice guitarParams={gParams} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最大価格</Th>
                        <Td><SearchMaxPrice guitarParams={gParams} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ソート</Th>
                        <Td><SelectorSort guitarParams={gParams} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>並び順</Th>
                        <Td><SelectorOrder guitarParams={gParams} /></Td>
                    </CommonBorderTr>
                    <CommonBorderTr styleObj={{borderBottom: "none"}}>
                        <Th>検索方法</Th>
                        <td style={{fontSize: "13px", paddingLeft: "20px"}}>
                            ※自動検索<br/>検索条件を変更すると<br/>自動的に検索されます。
                        </td>
                    </CommonBorderTr>
                </tbody>
            </table>
        </div>
    );
}

export default SearchConditionsGuitarMobile;