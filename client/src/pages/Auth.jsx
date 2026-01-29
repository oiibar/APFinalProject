import { useState } from "react";
import { fakeLogin } from "../api/api.js";

function Auth() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        fakeLogin(email, password);
        window.location.href = "/";
    };

    return (
        <div className="min-h-screen flex items-center justify-center">
            <form
                onSubmit={handleSubmit}
                className="bg-white p-8 rounded-2xl shadow-lg w-80"
            >
                <h1 className="text-2xl font-bold mb-6 text-center">Bookstore Login</h1>
                <input
                    className="w-full mb-4 p-2 border rounded"
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    className="w-full mb-4 p-2 border rounded"
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button className="w-full bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700">
                    Login
                </button>
            </form>
        </div>
    );
}

export default Auth;