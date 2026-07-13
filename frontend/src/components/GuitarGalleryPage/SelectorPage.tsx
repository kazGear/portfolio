import { useEffect } from "react";
import { GuitarParams, GuitarsResponse } from "../../types/Guitar";
import CommonPagination from "../common/CommonPagination";

interface ArgProps {
    guitarRes:    GuitarsResponse | null;
    guitarParams: GuitarParams;
    callback:     (gParams: GuitarParams) => Promise<void>;
}

const SelectorPage = ({guitarRes, guitarParams, callback}: ArgProps) => {
    const res     = guitarRes;
    const gParams = guitarParams;

    const changePrevPageHandler = () => {
        gParams.setPage(gParams.page - 1);
    }

    const changeNextPageHandler = () => {
        gParams.setPage(gParams.page + 1);
    }

    // ページを変更した時点で検索実行
    useEffect(() => {
        callback(gParams)
    }, [gParams.page])

    return (
        <CommonPagination changePrevPageHandler={changePrevPageHandler}
                          changeNextPageHandler={changeNextPageHandler}
                          hasPrev={res !== null ? res.hasPrev : false}
                          hasNext={res !== null ? res.hasNext : false}
                          styleObj={{margin: "-2px 0px 0px 20px"}}>
            <span> {res?.page} / {res?.totalPages} </span>
        </CommonPagination>
    );
}
export default SelectorPage;