import styled from "styled-components";
import { COLORS } from "../../lib/Constants";

const H3 = styled.h3`
    color: ${COLORS.ACCENT_FONT_GREEN};
    padding: 15px 40px;
`;

const TitleFrame = styled.div`
    background: ${COLORS.DARK_BACKGROUND};
    height: 100%;
    border-radius: 40px 0px 40px 0px;
`;

const Span = styled.span`
    color: ${COLORS.CAPTION_FONT};
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
                            <TitleFrame>
                                <H3># {idx + 1}&emsp;{c.historyTitle}</H3>
                            </TitleFrame>
                            <p><Span>期間：</Span>{c.period}</p>
                            <p><Span>業種：</Span>{c.industry}</p>
                            <p><Span>規模：</Span>{c.scale}</p>
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