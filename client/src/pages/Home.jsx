import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchBooks } from "../api/api";
import { useCart } from "../context/CartContext";

function Home() {
    const [books, setBooks] = useState([]);
    const { addItem } = useCart();

    useEffect(() => {
        (async () => {
            const b = await fetchBooks();
            setBooks(b);
        })();
    }, []);

    return (
        <div className="p-6">
            <div className="flex justify-between items-center mb-6 max-w-6xl mx-auto">
                <h1 className="text-3xl font-bold">Bookstore</h1>
                <div className="space-x-4">
                    <Link to="/cart" className="text-indigo-600 font-medium">
                        Cart
                    </Link>
                </div>
            </div>

            <div className="flex flex-wrap justify-center gap-6 max-w-6xl mx-auto">
                {books.map((book) => (
                    <div
                        key={book.id}
                        className="bg-white rounded-lg shadow-md p-4 w-60 flex flex-col"
                    >
                        <img
                            src={book.image}
                            alt={book.title}
                            className="w-full h-40 object-cover rounded"
                        />
                        <h2 className="text-lg font-semibold mt-2">{book.title}</h2>
                        <p className="text-sm text-gray-600">{book.author}</p>
                        <p className="font-bold mt-1">${book.price}</p>

                        {/* Push button to bottom */}
                        <div className="mt-auto">
                            <Link to={`/product/${book.id}`} className="text-sm text-indigo-600">
                                Details
                            </Link>
                            <button
                                onClick={() => addItem(book, 1)}
                                className="mt-2 w-full bg-indigo-500 text-white py-1.5 rounded hover:bg-indigo-600 transition"
                            >
                                Add to Cart
                            </button>
                        </div>
                    </div>

                ))}
            </div>
        </div>
    );
}

export default Home;
