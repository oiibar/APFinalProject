import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

function Auth() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [name, setName] = useState("");
    const [isRegister, setIsRegister] = useState(false);
    const [error, setError] = useState(null);
    const { login, register } = useAuth();
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);
        try {
            if (isRegister) {
                await register(name, email, password);
            } else {
                await login(email, password);
            }
            navigate("/");
        } catch (e) {
            setError(e.message || "Auth failed");
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center">
            <form
                onSubmit={handleSubmit}
                className="bg-white p-8 rounded-2xl shadow-lg w-80"
            >
                <h1 className="text-2xl font-bold mb-6 text-center">{isRegister ? "Register" : "Login"}</h1>
                {isRegister && (
                    <input
                        className="w-full mb-4 p-2 border rounded"
                        type="text"
                        placeholder="Name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                )}
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
                {error && <div className="text-red-500 mb-2">{error}</div>}
                <button className="w-full bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700">
                    {isRegister ? "Create account" : "Login"}
                </button>

                <div className="text-center mt-4 text-sm">
                    <button type="button" onClick={() => setIsRegister(!isRegister)} className="text-indigo-600">
                        {isRegister ? "Have an account? Login" : "No account? Register"}
                    </button>
                </div>
            </form>
        </div>
    );
}

export default Auth;