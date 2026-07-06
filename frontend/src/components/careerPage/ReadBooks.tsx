interface ArgProps {
    readBooks: string[];
}

const ReadBooks = ({readBooks}: ArgProps) => {
    return (
        <>
            <h3>購入書籍</h3>
            <hr/>
            <div>
            {
                readBooks.map(book =>
                    <p key={book}>{book}</p>
                )
            }
            </div>
            <p>※ その他５０冊以上</p>
        </>
    );
}
export default ReadBooks;