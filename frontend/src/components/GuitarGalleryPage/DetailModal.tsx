import styled from "styled-components";
import { Guitar } from "../../types/Guitar";
import { COLORS } from "../../lib/Constants";

const Sbackground = styled.div`
    width: 100%;
    height: 100%;
    display: ;
    position: absolute;
    z-index: 1000;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: rgba(0, 0, 0, 0.7);
`;

const Smodal = styled.div`
    width: 80%;
    height: 80%;
    display: ;
    border-radius: 20px;
    position: absolute;
    z-index: 2000;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: ${COLORS.BASE_BACKGROUND};
    box-shadow:
        inset 0 4px 0 rgba(255,255,255,0.5),
        inset 0 -8px 40px rgba(0,0,0,0.8);
`;

interface ArgProps {
    guitar: Guitar | null;
}

const DetailModal = (/*{guitar}: ArgProps*/) => {
    return (
        <Sbackground>
            <Smodal>

            </Smodal>
        </Sbackground>
    );
}
export default DetailModal;