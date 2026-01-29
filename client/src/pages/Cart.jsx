import { Link } from "react-router-dom";


function Cart() {
    const cart = JSON.parse(localStorage.getItem("cart")) || [];

    const total = cart.reduce((sum, item) => sum + item.price, 0);

    return (
        <div className="p-6">
            <Link to="/" className="text-indigo-600">Back</Link>
            <h1 className="text-3xl font-bold my-4">Shopping Cart</h1>


            {cart.length === 0 ? (
                <p>Your cart is empty.</p>
            ) : (
                <div className="space-y-4">
                    {cart.map((item, idx) => (
                        <div key={idx} className="bg-white p-4 rounded shadow flex justify-between">
                            <span>{item.title}</span>
                            <span>${item.price}</span>
                        </div>
                    ))}


                    <div className="font-bold text-xl">Total: ${total.toFixed(2)}</div>
                </div>
            )}
        </div>
    );
}

export default Cart;