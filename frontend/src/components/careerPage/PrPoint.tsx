interface ArgProps {
    prPoint: {
        point:   string;
        comment: string;
    }[];
}

const PrPoint = ({prPoint}: ArgProps) => {
    return (
        <>
            <h3>PR</h3>
            <hr/>
            <div>
            {
                prPoint.map((pr, idx) => {
                    return (
                        <>
                            <p key={idx + pr.point}>◆ {pr.point} ◆</p>
                            {
                                // 改行文字列が機能しないのでHTML<br/>に変換
                                pr.comment.split("<br/>").map(elem => {
                                    return (
                                        <>
                                            <span key={idx + elem}>{elem}</span>
                                            <br/>
                                        </>
                                    )
                                })
                            }
                        </>
                    )
                })
            }
            </div>
        </>
    );
}
export default PrPoint;