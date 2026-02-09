/* eslint-disable react-refresh/only-export-components */
import React, { createContext, useContext, useEffect, useState } from "react";
import { getCart as apiGetCart, updateCart as apiUpdateCart } from "../api/api";

const CartContext = createContext(null);

export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
    const [items, setItems] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        (async () => {
            const c = await apiGetCart();
            setItems(c);
            setLoading(false);
        })();
    }, []);

    const addItem = async (book, quantity = 1) => {
        const existing = items.find((it) => it.id === book.id);
        let updated;
        if (existing) {
            updated = items.map((it) =>
                it.id === book.id ? { ...it, quantity: Math.min((it.quantity || 1) + quantity, book.stock) } : it
            );
        } else {
            updated = [...items, { ...book, quantity }];
        }
        setItems(updated);
        await apiUpdateCart(updated);
    };

    const removeItem = async (bookId) => {
        const updated = items.filter((it) => it.id !== bookId);
        setItems(updated);
        await apiUpdateCart(updated);
    };

    const updateQuantity = async (bookId, quantity) => {
        const updated = items.map((it) => (it.id === bookId ? { ...it, quantity } : it));
        setItems(updated);
        await apiUpdateCart(updated);
    };

    const clear = async () => {
        setItems([]);
        await apiUpdateCart([]);
    };

    const subtotal = items.reduce((s, it) => s + it.price * (it.quantity || 1), 0);

    return (
        <CartContext.Provider value={{ items, loading, addItem, removeItem, updateQuantity, clear, subtotal }}>
            {children}
        </CartContext.Provider>
    );
};
