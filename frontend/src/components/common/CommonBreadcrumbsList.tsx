import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Checkbox = styled.input`
    display: inline-block;
    border: 3px solid ${COLORS.BORDER};
    border-radius: 5px;
    padding: 4px;
    margin: 5px;
    background: ${COLORS.BASE_BACKGROUND};
`

interface ArgProps {
    children:  React.ReactNode;
    value:     string | number;
    onChange?: React.ChangeEventHandler<HTMLInputElement>;
}

const CommonBreadcrumbsList = ({children, value, onChange}: ArgProps) => {
    return (
        <label style={{display: "inline-block"}}>
            <Checkbox type="checkbox" value={value} onChange={onChange} />
            <span style={{display: "inline-block", transform: "translateY(-2px)"}}>
                {children}
            </span>
        </label>
    );
}

export default CommonBreadcrumbsList;