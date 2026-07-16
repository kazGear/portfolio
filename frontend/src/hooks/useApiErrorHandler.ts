import { useNavigate } from "react-router-dom";
import { ApiError } from "../types/ApiError";

const useApiErrorHandler = () => {
    const navigate = useNavigate();

    return (error: unknown) => {
        if (error instanceof ApiError) {
            if (error.status >= 500) {
                navigate("/ErrorPage");
            }
        }
    };
}

export default useApiErrorHandler;