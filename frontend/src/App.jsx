import './App.css';
import './animation.css';
import IndexPage from './pages/IndexPage';
import { Route, Routes, useNavigate } from 'react-router-dom';
import BattlePage from './pages/BattlePage';
import ShopPage from './pages/ShopPage';
import LoginPage from "./pages/LoginPage";
import CommonAppHeader from './components/common/CommonAppHeader';
import BattleResultPage from './pages/BattleResultPage';
import UserPage from "./pages/UserPage";
import EditPage from "./pages/EditPage";
import GuitarGalleryPage from "./pages/GuitarGalleryPage";
import CareerPage from "./pages/CareerPage";
import ErrorPage from "./pages/ErrorPage";
import CommonErrorBoundary from "./components/common/CommonErrorBoundary";
import { SIZE } from "./lib/Constants";

function App() {

    return (
        <CommonErrorBoundary>
            <CommonAppHeader title="KazApp" />
            <main style={{paddingTop: SIZE.HEADER_HEIGHT}}>
                <Routes>
                    {/* 新しいページを作成したらここに追加（要:import） */}
                    <Route path={"/"} element={<IndexPage />} />
                    <Route path={"/IndexPage"} element={<IndexPage />} />
                    <Route path={"/LoginPage"} element={<LoginPage />} />
                    <Route path={"/ShopPage"} element={<ShopPage />} />
                    <Route path={"/BattlePage"} element={<BattlePage />} />
                    <Route path={"/BattleResultPage"} element={<BattleResultPage />} />
                    <Route path={"/UserPage"} element={<UserPage />} />
                    <Route path={"/EditPage"} element={<EditPage />} />
                    <Route path={"/GuitarGalleryPage"} element={<GuitarGalleryPage />} />
                    <Route path={"/CareerPage"} element={<CareerPage />} />
                    <Route path={"/ErrorPage"} element={<ErrorPage />} />
                </Routes>
            </main>
        </CommonErrorBoundary>
    );
}

export default App;
