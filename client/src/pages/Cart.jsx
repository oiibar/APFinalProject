import {Link, useNavigate} from "react-router";
import {useCart} from "../context/CartContext";


function Cart() {
    const {items, subtotal, updateQuantity, removeItem} = useCart();
    const navigate = useNavigate();

    const changeNumber = (id, delta) => {
        const it = items.find((x) => x.id === id);
        if (!it) return;
        const next = Math.max(1, (it.quantity || 1) + delta);
        const final = Math.min(next, it.stock || 9999);
        updateQuantity(id, final);
    };

    return (
        <div className="p-6 max-w-6xl mx-auto">
            <Link to="/" className="text-indigo-600">Back</Link>
            <h1 className="text-3xl font-bold my-4">Shopping Cart</h1>

            {items.length === 0 ? (
                <p>Your cart is empty.</p>
            ) : (
                <div className="space-y-4 max-w-3xl">
                    {items.map((item) => (
                        <div key={item.id} className="bg-white p-4 rounded shadow flex justify-between items-center">
                            <div>
                                <div className="font-semibold">{item.title}</div>
                                <div className="text-sm text-gray-600">{item.author}</div>
                            </div>

                            <div className="flex items-center gap-4">
                                <div className="flex items-center border rounded">
                                    <button onClick={() => changeNumber(item.id, -1)} className="px-3">-</button>
                                    <div className="px-3">{item.quantity || 1}</div>
                                    <button onClick={() => changeNumber(item.id, 1)} className="px-3">+</button>
                                </div>

                                <div className="w-24 text-right">${(item.price * (item.quantity || 1)).toFixed(2)}</div>
                                <button onClick={() => removeItem(item.id)} className="text-red-500">Remove</button>
                            </div>
                        </div>
                    ))}


                    <div className="font-bold text-xl">Total: ${subtotal.toFixed(2)}</div>
                    <div className="flex gap-3">
                        <button onClick={() => navigate('/checkout')}
                                className="bg-indigo-600 text-white px-4 py-2 rounded">Proceed to Checkout
                        </button>
                        <Link to="/" className="px-4 py-2 border rounded">Continue shopping</Link>
                    </div>
                </div>
            )}
        </div>
    );
}

export default Cart;