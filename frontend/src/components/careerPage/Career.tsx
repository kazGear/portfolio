import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const Sh3 = styled.h3`
    color: ${COLORS.ACCENT_FONT_GREEN};
    padding: 15px 40px;
`;

const StitleFrame = styled.div`
    background: ${COLORS.DARK_BACKGROUND};
    height: 100%;
    border-radius: 40px 0px 40px 0px;
`;

const Sspan = styled.span`
    color: ${COLORS.CAPTION_FONT_COLOR};
    font-weight: bold;
`;

interface ArgProps {
    career: {
        historyTitle:         string;
        period:               string;
        industry:             string;
        scale:                string;
        programmingLanguages: string;
        jobContents:          string[];
    }[]
}

const Career = ({career}: ArgProps) => {
    return (
        <>
            <h3>職務経歴</h3>
            {
                career.map((c, idx) => {
                    return (
                        <div key={c.historyTitle}>
                            <StitleFrame>
                                <Sh3># {idx + 1}&emsp;{c.historyTitle}</Sh3>
                            </StitleFrame>
                            <p><Sspan>期間：</Sspan>{c.period}</p>
                            <p><Sspan>業種：</Sspan>{c.industry}</p>
                            <p><Sspan>規模：</Sspan>{c.scale}</p>
                            {
                                c.jobContents.map(content => <p key={content}>{content}</p>)
                            }
                        </div>
                    )
                })
            }
        </>
    );
}
export default Career;