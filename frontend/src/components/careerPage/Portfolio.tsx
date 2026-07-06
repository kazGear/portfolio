import { COLORS } from "../../lib/Constants";

interface ArgProps {
    portfolio: {
        appURLCaption:    string;
        appURL:           string;
        sourceURLCaption: string;
        sourceURL:        string;
        comment:          string[];
    }
}

const Portfolio = ({portfolio}: ArgProps) => {
    return (
        <>
            <h3>ポートフォリオ</h3>
            <p>{portfolio.appURLCaption}</p>

            <div>
                <a href={portfolio.appURL} target="_blank" style={{color: COLORS.URL}}>
                    {portfolio.appURL}
                </a>
            </div>

            <p>{portfolio.sourceURLCaption}</p>

            <div>
                <a target="_blank" href={portfolio.sourceURL} style={{color: COLORS.URL}}>
                    {portfolio.sourceURL}
                </a>
            </div>

            <div>
            {
                portfolio.comment.map(
                    (comment, idx) => {
                        return comment !== "<br/>" ? <p key={comment}>{comment}</p>
                                                   : <br key={idx + `${comment}`}/>
                    }
                )
            }
            </div>
        </>
    );
}
export default Portfolio;