import { GuitarParams } from "../../types/Guitar";
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

const Sp = styled.p`
    font-size: 13px;
    font-weight: bolder;
    margin-top: 14.5px;
`;

interface ArgProps {
    guitarParams:  GuitarParams;
    makers:        Code[] | null;
    series:        Code[] | null;
    colors:        Code[] | null;
    bodyMaterials: Code[] | null;
    callback: (gParams: GuitarParams) => Promise<void>;
}

const SearchConditions = ({guitarParams,
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
                <Sp>メーカー</Sp>
                <Sp>カラー</Sp>
                <Sp>シリーズ</Sp>
                <Sp>ギター名</Sp>
                <Sp>トップ材</Sp>
                <Sp>ボディ材</Sp>
                <Sp>最低価格</Sp>
                <Sp>最大価格</Sp>
                <Sp>ソート</Sp>
                <Sp>並び順</Sp>
                <Sp>選択ページ</Sp>
                <Sp>ページサイズ</Sp>
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
                <SelectorPage guitarParams={gParams} callback={callback}/>
                <SelectorPageSize guitarParams={gParams} callback={callback}/>
                <p style={{marginRight: "20px", textAlign: "left"}}>
                    ※自動検索<br/>検索条件を変更すると自動的に検索されます。
                </p>
            </div>
        </div>
    );
}
export default SearchConditions;