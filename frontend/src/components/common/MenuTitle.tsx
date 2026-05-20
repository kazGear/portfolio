import styled from "styled-components";

const SmenuTitle = styled.h2`
    height: 40px;
    display: flex;
    align-items: center;
    padding-left: 60px;
    border-radius: 10px;
    border-collapse: collapse;
`;

interface ArgProps {
    title: string
    styleObj?: React.CSSProperties;
    className?: string
}

const MenuTitle = ({title, styleObj, className}: ArgProps) => {
    return <SmenuTitle className={className} style={styleObj}>
               {title}
           </SmenuTitle>;
};

export default MenuTitle;
