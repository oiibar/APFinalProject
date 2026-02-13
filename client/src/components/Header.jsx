import React from "react";
import { Link } from "react-router";
import { useAuth } from "../context/AuthContext";
import { useCart } from "../context/CartContext";

export default function Header() {
    const { user, logout } = useAuth();
    const { items, subtotal } = useCart();

    return (
        <header className="bg-white shadow-sm py-3 mb-6">
            <div className="max-w-6xl mx-auto px-4 flex justify-between items-center">
                <Link to="/" className="text-2xl font-bold">
                    Bookstore
                </Link>

                <nav className="flex items-center space-x-4">
                    <Link to="/" className="text-gray-700 hover:text-indigo-600">
                        Home
                    </Link>
                    <Link to="/cart" className="text-gray-700 hover:text-indigo-600">
                        Cart ({items.length})
                    </Link>
                    <div className="text-sm text-gray-600">${subtotal.toFixed(2)}</div>

                    {user ? (
                        <div className="flex items-center space-x-3">
                            <span className="text-gray-700">{user.name}</span>
                            <button
                                onClick={async () => {
                                    await logout();
                                    window.location.href = "/login";
                                }}
                                className="text-sm text-red-500"
                            >
                                Logout
                            </button>
                        </div>
                    ) : (
                        <Link to="/login" className="text-indigo-600 font-medium">
                            Login
                        </Link>
                    )}
                </nav>
            </div>
        </header>
    );
}
