import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Auth from "./pages/Auth";
import Home from "./pages/Home";
import Cart from "./pages/Cart";
import Product from "./pages/Product";
import Checkout from "./pages/Checkout";
import { isAuthenticated } from "./api/api.js";
import { AuthProvider } from "./context/AuthContext";
import { CartProvider } from "./context/CartContext";
import Header from "./components/Header";

function App() {
    return (
        <AuthProvider>
            <CartProvider>
                <BrowserRouter>
                    <Header />
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
                        <Route
                            path="/product/:id"
                            element={isAuthenticated() ? <Product /> : <Navigate to="/login" />}
                        />
                        <Route
                            path="/checkout"
                            element={isAuthenticated() ? <Checkout /> : <Navigate to="/login" />}
                        />
                    </Routes>
                </BrowserRouter>
            </CartProvider>
        </AuthProvider>
    );
}

export default App;