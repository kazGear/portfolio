import { GuitarParams, GuitarsResponse } from "../../types/Guitar";
import { Code } from "../../types/Code";
import SearchMaker from "./SearchMaker";
import SearchColor from "./SearchColor";
import SearchSeries from "./SearchSeries";
import SearchName from "./SearchName";
import SearchMaterialTop from "./SearchMaterialTop";
import SearchMaterialBack from "./SearchMaterialBack";
import SearchMinPrice from "./SearchMinPrice";
import SearchMaxPrice from "./SearchMaxPrice";
import styled from "styled-components";
import SelectorOrder from "./SelectorOrder";
import SelectorSort from "./SelectorSort";
import SelectorPage from "./SelectorPage";
import SelectorPageSize from "./SelectorPageSize";
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
    guitarRes:     GuitarsResponse | null;
    guitarParams:  GuitarParams;
    makers:        Code[] | null;
    series:        Code[] | null;
    colors:        Code[] | null;
    bodyMaterials: Code[] | null;
    callback:      (gParams: GuitarParams) => Promise<void>;
}

const SearchConditions = ({guitarRes,
                           guitarParams,
                           makers,
                           series,
                           colors,
                           bodyMaterials,
                           callback}: ArgProps
) => {
    const gParams = guitarParams;

    return (
        <div style={{margin: "10px", overflow: "hidden"}}>
            <table>
                <thead>
                    <CommonBorderTr>
                        <th>検索条件</th>
                        <td style={{paddingLeft: "30px"}}>設定値</td>
                    </CommonBorderTr>
                </thead>
                <tbody>
                    <CommonBorderTr>
                        <Th>メーカー</Th>
                        <Td><SearchMaker guitarParams={gParams} makers={makers} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>カラー</Th>
                        <Td><SearchColor guitarParams={gParams} colors={colors} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>シリーズ</Th>
                        <Td><SearchSeries guitarParams={gParams} series={series} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ギター名</Th>
                        <Td><SearchName guitarParams={gParams} callback={callback} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>トップ材</Th>
                        <Td><SearchMaterialTop guitarParams={gParams} materials={bodyMaterials} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ボディ材</Th>
                        <Td><SearchMaterialBack guitarParams={gParams} materials={bodyMaterials} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最低価格</Th>
                        <Td><SearchMinPrice guitarParams={gParams} callback={callback} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>最大価格</Th>
                        <Td><SearchMaxPrice guitarParams={gParams} callback={callback} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ソート</Th>
                        <Td><SelectorSort guitarParams={gParams} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>並び順</Th>
                        <Td><SelectorOrder guitarParams={gParams} callback={callback}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>ページサイズ</Th>
                        <Td><SelectorPageSize guitarParams={gParams} callback={callback} styleObj={styleObj}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>選択ページ</Th>
                        <Td><SelectorPage guitarParams={gParams}
                                          guitarRes={guitarRes}
                                          callback={callback}
                                          styleObj={{margin: "0px 0px 6px 15px"}}/></Td>
                    </CommonBorderTr>
                    <CommonBorderTr>
                        <Th>検索方法</Th>
                        <Td>※自動検索<br/>検索条件を変更すると<br/>自動的に検索されます。</Td>
                    </CommonBorderTr>
                </tbody>
            </table>
        </div>
    );
}

export default SearchConditions;