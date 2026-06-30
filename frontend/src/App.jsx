import './App.css';
import './animation.css';
import IndexPage from './pages/IndexPage';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import BattlePage from './pages/BattlePage';
import ShopPage from './pages/ShopPage';
import LoginPage from "./pages/LoginPage";
import AppHeader from './components/common/AppHeader';
import BattleResultPage from './pages/BattleResultPage';
import UserPage from "./pages/UserPage";
import EditPage from "./pages/EditPage";
import GuitarGalleryPage from "./pages/GuitarGalleryPage";
import ErrorBoundary from "./components/common/ErrorBoundary";

function App() {
    return (
        <BrowserRouter>
            <ErrorBoundary>
                <AppHeader title="KazApp" />
                <main style={{paddingTop:"60px"}}>
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
                    </Routes>
                </main>
            </ErrorBoundary>
        </BrowserRouter>
    );
}

export default App;
