import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { fetchBookById } from "../api/api";
import { useCart } from "../context/CartContext";

export default function Product() {
    const { id } = useParams();
    const [book, setBook] = useState(null);
    const { addItem } = useCart();

    useEffect(() => {
        (async () => {
            const b = await fetchBookById(id);
            setBook(b);
        })();
    }, [id]);

    if (!book) return <div className="p-6">Loading...</div>;

    return (
        <div className="p-6 max-w-4xl mx-auto">
            <Link to="/" className="text-indigo-600">Back</Link>
            <div className="flex gap-6 mt-4 items-start">
                <img src={book.image} alt={book.title} className="w-72 h-96 object-cover rounded" />
                <div>
                    <h1 className="text-2xl font-bold">{book.title}</h1>
                    <p className="text-gray-600">by {book.author}</p>
                    <p className="mt-4">{book.description}</p>
                    <div className="mt-4">
                        <div>Pages: {book.pages}</div>
                        <div>Publisher: {book.publisher}</div>
                        <div>ISBN: {book.isbn}</div>
                        <div>In stock: {book.stock}</div>
                        <div className="text-2xl font-bold mt-2">${book.price}</div>
                    </div>

                    <div className="mt-4 flex items-center gap-3">
                        <button
                            onClick={() => addItem(book, 1)}
                            className="bg-indigo-600 text-white px-4 py-2 rounded"
                        >
                            Add to cart
                        </button>
                        <Link to="/cart" className="text-indigo-600">Go to cart</Link>
                    </div>
                </div>
            </div>
        </div>
    );
}

