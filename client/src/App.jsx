import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Auth from "./pages/Auth";
import Home from "./pages/Home";
import Cart from "./pages/Cart";
import { isAuthenticated } from "./api/api.js";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Auth />} />
                <Route
                    path="/"
                    element={isAuthenticated() ? <Home /> : <Navigate to="/login" />}
                />
                <Route
                    path="/cart"
                    element={isAuthenticated() ? <Cart /> : <Navigate to="/login" />}
                />
            </Routes>
        </BrowserRouter>
    );
}

export default App;