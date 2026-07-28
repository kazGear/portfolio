import { useState } from "react";

type Position = {
    x: number;
    y: number;
};

interface ArgProps {
    imgURL: string | undefined;
    alt:    string | undefined;
    width:  number;
    height: number;
    zoomRate: number;
}

const CommonZoomableImage = ({ imgURL, alt,  width, height, zoomRate }: ArgProps) => {
    const initPosition: Position  = { x: -1000, y: -1000}
    const [position, setPosition] = useState<Position>(initPosition);
    const [isShow, setIsShow]     = useState<boolean>(true);

    // 画像拡大位置の初期化（拡大画像を表示しない）
    const mouseLeaveHandler = () => {
        setPosition(initPosition);
        setIsShow(true);
    }

    // 拡大中に元画像は不要
    const mouseEnterHandler = () => setIsShow(false);

    // 画像拡大の位置を調整
    const mouseMoveHandler = (event: React.MouseEvent<HTMLImageElement>) => {
        // 画像情報を取得
        const rect = event.currentTarget.getBoundingClientRect();

        // 画像上でのポインタ位置を算出
        const mouseX = event.clientX - rect.left;
        const mouseY = event.clientY - rect.top;

        // %単位に変換し、background-position(拡大画像表示位置)で使用
        const x = (mouseX / rect.width) * 100;
        const y = (mouseY / rect.height) * 100;

        setPosition({ x, y });
    };

    return (
        <div style={{
            textAlign: "center",
            margin: "auto",
            backgroundImage: `url(${imgURL})`,
            backgroundRepeat: "no-repeat",
            backgroundSize: `${zoomRate}%`,
            backgroundPosition: `${position.x}% ${position.y}%`,
        }} id="zoom-test">
            <img
                src={imgURL}
                alt={alt}
                onMouseEnter={mouseEnterHandler}
                onMouseMove={mouseMoveHandler}
                onMouseLeave={mouseLeaveHandler}
                style={{
                    opacity: `${isShow ? 1.0 : 0.0}`,
                    width: `${width}px`,
                    height: `${height}px`,
                    objectFit: "contain",
            }}/>
        </div>
    );
};

export default CommonZoomableImage;