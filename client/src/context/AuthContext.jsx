/* eslint-disable react-refresh/only-export-components */
import React, { createContext, useContext, useState } from "react";
import { getCurrentUser, login as apiLogin, register as apiRegister, logout as apiLogout } from "../api/api";

const AuthContext = createContext(null);

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(getCurrentUser());
    const [loading, setLoading] = useState(false);

    const login = async (email, password) => {
        setLoading(true);
        try {
            const u = await apiLogin({ email, password });
            setUser(u);
            setLoading(false);
            return u;
        } catch (err) {
            setLoading(false);
            throw err;
        }
    };

    const register = async (name, email, password) => {
        setLoading(true);
        try {
            const u = await apiRegister({ name, email, password });
            setUser(u);
            setLoading(false);
            return u;
        } catch (err) {
            setLoading(false);
            throw err;
        }
    };

    const logout = async () => {
        setLoading(true);
        await apiLogout();
        setUser(null);
        setLoading(false);
    };

    return (
        <AuthContext.Provider value={{ user, loading, login, register, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
