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
import Button from "../common/Button";

const Sp = styled.p`
    font-size: 13px;
    font-weight: bolder;
    margin-top: 14.5px;
`;

interface ArgProps {
    guitarParams:        GuitarParams;
    makers:              Code[] | null;
    series:              Code[] | null;
    colors:              Code[] | null;
    bodyMaterials:       Code[] | null;
    guitarSearchHandler: (gParams: GuitarParams) => Promise<void>;
}

const SearchConditions = ({guitarParams,
                           makers,
                           series,
                           colors,
                           bodyMaterials,
                           guitarSearchHandler}: ArgProps
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
                <SearchMaker guitarParams={gParams}
                             makers={makers}/>
                <SearchColor guitarParams={gParams}
                             colors={colors}/>
                <SearchSeries guitarParams={gParams}
                              series={series}/>
                <SearchName guitarParams={gParams}/>
                <SearchMaterialTop guitarParams={gParams}
                                   materials={bodyMaterials}/>
                <SearchMaterialBack guitarParams={gParams}
                                    materials={bodyMaterials}/>
                <SearchMinPrice guitarParams={gParams}/>
                <SearchMaxPrice guitarParams={gParams}/>
                <SelectorSort guitarParams={gParams}/>
                <SelectorOrder guitarParams={gParams}/>
                <SelectorPage guitarParams={gParams}/>
                <SelectorPageSize guitarParams={gParams}/>
                <Button text="検索"
                        onClick={() => guitarSearchHandler(gParams)}
                        styleObj={{margin: "15px 20px 0px 0px"}}>

                </Button>
            </div>
        </div>
    );
}
export default SearchConditions;