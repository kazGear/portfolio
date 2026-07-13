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

const P = styled.p`
    font-size: 13px;
    font-weight: bolder;
    margin-top: 14.5px;
`;

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
        <div style={{display: "flex", margin: "10px"}}>
            <div style={{marginTop: "10px", minWidth: "80px"}}>
                {/* ラベル、インプットの並び順を合わせること
                    TODO: スタイルで無理やりラベルとインプットの高さを揃えているため、綺麗に行として揃える。
                          Input, Selectコンポーネントは内部でdivで囲っているため使いづらい。これもリファクタ対象。
                */}
                <P>メーカー</P>
                <P>カラー</P>
                <P>シリーズ</P>
                <P>ギター名</P>
                <P>トップ材</P>
                <P>ボディ材</P>
                <P>最低価格</P>
                <P>最大価格</P>
                <P>ソート</P>
                <P>並び順</P>
                <P>選択ページ</P>
                <P>ページサイズ</P>
            </div>
            <div style={{marginTop: "10px", textAlign: "right"}}>
                <SearchMaker guitarParams={gParams} makers={makers} callback={callback}/>
                <SearchColor guitarParams={gParams} colors={colors} callback={callback}/>
                <SearchSeries guitarParams={gParams} series={series} callback={callback}/>
                <SearchName guitarParams={gParams} callback={callback}/>
                <SearchMaterialTop guitarParams={gParams} materials={bodyMaterials} callback={callback}/>
                <SearchMaterialBack guitarParams={gParams} materials={bodyMaterials} callback={callback}/>
                <SearchMinPrice guitarParams={gParams} callback={callback}/>
                <SearchMaxPrice guitarParams={gParams} callback={callback}/>
                <SelectorSort guitarParams={gParams} callback={callback}/>
                <SelectorOrder guitarParams={gParams} callback={callback}/>
                <SelectorPage guitarParams={gParams} guitarRes={guitarRes} callback={callback}/>
                <SelectorPageSize guitarParams={gParams} callback={callback}/>
                <p style={{marginRight: "20px", textAlign: "left"}}>
                    ※自動検索<br/>検索条件を変更すると自動的に検索されます。
                </p>
            </div>
        </div>
    );
}
export default SearchConditions;