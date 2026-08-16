import styled from "styled-components";
import { COLORS } from "../../lib/Constants";
import { Job } from "../../types/Job";
import { Link } from "react-router-dom";

const CardFrame = styled.div`
    font-weight: 900;
    width:  100%;
    margin: 20px 0px;
    padding: 10px 0px;
    background: ${COLORS.BASE_BACKGROUND};
    border-radius: 15px;
    box-shadow:
        inset 0 1px 0 rgba(255,255,255,0.1),
        inset 0 -2px 10px rgba(0,0,0,0.4);
`;

const Span = styled.span`
    display: inline-block;
    border: 3px solid ${COLORS.BORDER};
    border-radius: 5px;
    padding: 4px;
    margin: 5px;
    background: ${COLORS.BASE_BACKGROUND};
`

const Th = styled.th`
    padding-right: 10px;
    min-width: 80px;
`

const P = styled.p`
    margin: 10px;
`

interface ArgProps {
    job: Job | null;
}

const JobCard = ({ job }: ArgProps) => {

    return (
        <Link to={job!.Url} target="_blank" style={{textDecoration: "none", color: "inherit"}}>
            <CardFrame>
                <table style={{margin: "20px"}}>
                    <thead>
                        <tr>
                            <Th>--項目--</Th>
                            <td style={{textAlign: "center"}}>----内容----</td>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <Th>タイトル</Th>
                            <td><h3 style={{color: COLORS.ACCENT_FONT_PINK}}>{job?.Title}</h3></td>
                        </tr>
                        <tr>
                            <Th>所在地</Th>
                            <td><P>{job?.Location}</P></td>
                        </tr>
                        <tr>
                            <Th>勤務形態</Th>
                            <td><P>{job?.WorkPlace}</P></td>
                        </tr>
                        <tr>
                            <Th>報酬</Th>
                            <td><P>{job?.MinSalaryAtMonth}&nbsp;～&nbsp;{job?.MaxSalaryAtMonth}&nbsp;円</P></td>
                        </tr>
                        <tr>
                            <Th>特徴</Th>
                            <td>
                                {
                                    job?.FeatureNames.map(
                                        feature => <Span key={job.Url + feature}>{feature}</Span>
                                    )
                                }
                            </td>
                        </tr>
                        <tr>
                            <Th>オプション</Th>
                            <td>
                                {
                                    job?.Options.map(
                                        option => <Span key={job.Url + option}>{option}</Span>
                                    )
                                }
                            </td>
                        </tr>
                        <tr>
                            <Th>ソース</Th>
                            <td><P>{job?.SourceSite}</P></td>
                        </tr>
                    </tbody>
                </table>
            </CardFrame>
        </Link>
    );
}
export default JobCard;