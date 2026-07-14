import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import React from "react";
import { MetaDataDTO } from "../../types/MonsterBattle";

const MessageFrame = styled.div`
    min-height: 30vh;
    border: inset 4px ${COLORS.BORDER};
    overflow-y: scroll;
    background: rgba(255, 255, 255, 0.8);
`;
const MessageWindow = styled.div`
    margin: 10px;
`;

interface ArgProps { shortLog: MetaDataDTO[]; }

const MessageWindowBlock = ({shortLog}: ArgProps) => {
    return (
        <MessageFrame id="messageWindow">
            <MessageWindow>
                {
                    shortLog.map((log, index) => {
                        return (
                            <React.Fragment key={index + log.Message}>
                                {log.Message}<br />
                            </React.Fragment>
                        )
                    })
                }
            </MessageWindow>
        </MessageFrame>
    );
}

export default MessageWindowBlock;