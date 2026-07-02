import { GuitarsResponse } from "../../types/Guitar";
import GuitarCard from "./GuitarCard";

interface ArgProps {
    // children: React.ReactNode;
    guitarsRes: GuitarsResponse | null;
}

const GuitarCards = ({guitarsRes: res}: ArgProps) => {
    return (
        <>
            <p style={{marginLeft: "15px", fontWeight: "bolder"}}>
                検索結果 {res?.totalCount} 件<br/>
                ページ {res?.page} / {res?.totalPages} @{res?.pageSize}件
            </p>
            <div style={{display: "flex", flexWrap: "wrap", justifyContent: "space-evenly", margin: "15px"}}>
                {
                    res?.guitars.map(guitar => (
                        <GuitarCard guitar={guitar}
                                    key={guitar.maker + guitar.name + guitar.color} />
                    )
                )}
            </div>
        </>
    );
}

export default GuitarCards;