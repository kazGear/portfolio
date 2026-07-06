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
                            <p key={pr.point}>◆ {pr.point}</p>
                            <p key={pr.point + "comment"}>{pr.comment}</p>
                        </>
                    )
                })
            }
            </div>
        </>
    );
}
export default PrPoint;