interface ArgProps {
    specialty: string[];
}

const Specialty = ({specialty}: ArgProps) => {
    return (
        <>
            <h3>得意分野</h3>
            <hr/>
            <div>
            {
                specialty.map(s =>
                    <p key={s}>・{s}</p>
                )
            }
            </div>
        </>
    );
}
export default Specialty;