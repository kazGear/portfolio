import { GuitarsResponse } from "../../types/Guitar";
import GuitarCard from "./GuitarCard";

interface ArgProps {
    // children: React.ReactNode;
    guitarsRes: GuitarsResponse | null;
}

const GuitarCards = ({guitarsRes: guitarsResponse}: ArgProps) => {
    return (
        <div style={{display: "flex", flexWrap: "wrap", justifyContent: "space-evenly", margin: "15px"}}>
            {
                guitarsResponse?.guitars.map(guitar => (
                    <GuitarCard guitar={guitar}
                                key={guitar.maker + guitar.name + guitar.color} />
                )
            )}
        </div>
    );
}

export default GuitarCards;