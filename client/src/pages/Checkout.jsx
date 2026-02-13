import { useState } from "react";
import { useCart } from "../context/CartContext";
import { useAuth } from "../context/AuthContext";
import { createOrder } from "../api/api";

export default function Checkout() {
    const { items, subtotal, clear } = useCart();
    const { user } = useAuth();
    const [address, setAddress] = useState("");
    const [loading, setLoading] = useState(false);
    const [order, setOrder] = useState(null);
    const [error, setError] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError(null);
        try {
            const result = await createOrder({ user, items, address });
            setOrder(result);
            await clear();
        } catch (e) {
            setError(e.message || "Failed to create order");
        } finally {
            setLoading(false);
        }
    };

    if (order)
        return (
            <div className="p-6 max-w-3xl mx-auto">
                <h1 className="text-2xl font-bold">Order confirmed</h1>
                <p className="mt-4">Thank you, your order id: <span className="font-mono">{order.orderId}</span></p>
            </div>
        );

    return (
        <div className="p-6 max-w-3xl mx-auto">
            <h1 className="text-2xl font-bold mb-4">Checkout</h1>
            <div className="mb-4">Items: {items.length} — Total: ${subtotal.toFixed(2)}</div>

            <form onSubmit={handleSubmit} className="space-y-3 bg-white p-4 rounded shadow">
                <div>
                    <label className="block text-sm text-gray-700">Email</label>
                    <input className="w-full border p-2 rounded" value={user?.email || ""} readOnly />
                </div>
                <div>
                    <label className="block text-sm text-gray-700">Shipping address</label>
                    <textarea required className="w-full border p-2 rounded" value={address} onChange={(e) => setAddress(e.target.value)} />
                </div>

                {error && <div className="text-red-500">{error}</div>}

                <button disabled={loading || items.length === 0} className="bg-indigo-600 text-white px-4 py-2 rounded">
                    {loading ? "Processing..." : "Place order"}
                </button>
            </form>
        </div>
    );
}

