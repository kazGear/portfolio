import './App.css';
import './animation.css';
import IndexPage from './pages/IndexPage';
import { Route, Routes } from 'react-router-dom';
import BattlePage from './pages/BattlePage';
import ShopPage from './pages/ShopPage';
import LoginPage from "./pages/LoginPage";
import CommonAppHeader from './components/common/CommonAppHeader';
import BattleResultPage from './pages/BattleResultPage';
import UserPage from "./pages/UserPage";
import EditPage from "./pages/EditPage";
import GuitarGalleryPage from "./pages/GuitarGalleryPage";
import GuitarGalleryDetailPage from "./pages/GuitarGalleryDetailPage";
import CareerPage from "./pages/CareerPage";
import JobPage from "./pages/JobPage";
import JobAnalyzePage from "./pages/JobAnalyzePage";
import ErrorPage from "./pages/ErrorPage";
import CommonErrorBoundary from "./components/common/CommonErrorBoundary";
import { COLORS, SIZE } from "./lib/Constants";
import styled from 'styled-components';

const Main = styled.main`
    padding-top: ${SIZE.HEADER_HEIGHT};

    @media (max-width: ${SIZE.MOBILE_LAYOUT_BREAKPOINT}) {
        padding-top: 0px;
    }
`;

function App() {

    return (
        <CommonErrorBoundary>
            <CommonAppHeader title="KazApp" />
            <Main style={{color: COLORS.MAIN_FONT}}>
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
                    <Route path={"/GuitarGalleryPage/:makerCd/:name/:color"} element={<GuitarGalleryDetailPage />} />
                    <Route path={"/CareerPage"} element={<CareerPage />} />
                    <Route path={"/JobPage"} element={<JobPage />} />
                    <Route path={"/JobAnalyzePage"} element={<JobAnalyzePage />} />
                    <Route path={"/ErrorPage"} element={<ErrorPage />} />
                </Routes>
            </Main>
        </CommonErrorBoundary>
    );
}

export default App;
