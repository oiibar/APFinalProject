// import axios from "axios";
//
// export const books = [
//     {
//         id: "book-1",
//         title: "Book 1",
//         author: "John Doe",
//         price: 9.99,
//         image: "https://placehold.co/70x70",
//         description: "Book 1 description.",
//         pages: 200,
//         publisher: "Publisher 1",
//         stock: 8,
//     },
//     {
//         id: "book-2",
//         title: "Book 2",
//         author: "John Doe",
//         price: 9.99,
//         image: "https://placehold.co/70x70",
//         description: "Book 2 description.",
//         pages: 300,
//         publisher: "Publisher 2",
//         stock: 5,
//     },
//     {
//         id: "book-3",
//         title: "Book 3",
//         author: "John Doe",
//         price: 9.99,
//         image: "https://placehold.co/70x70",
//         description: "Book 3 description.",
//         pages: 400,
//         publisher: "Publisher 3",
//         stock: 3,
//     },
//     {
//         id: "book-4",
//         title: "Book 4",
//         author: "John Doe",
//         price: 9.99,
//         image: "https://placehold.co/70x70",
//         description: "Book 4 description.",
//         pages: 500,
//         publisher: "Publisher 4",
//         stock: 10,
//     },
//     {
//         id: "book-6",
//         title: "Book 5",
//         author: "John Doe",
//         price: 9.99,
//         image: "https://placehold.co/70x70",
//         description: "Book 5 description.",
//         pages: 100,
//         publisher: "Publisher 5",
//         stock: 6,
//     },
// ];
//
// const getStoredUsers = () => {
//     try {
//         const raw = localStorage.getItem("user");
//         return raw ? JSON.parse(raw) : [{
//             id: "user-1",
//             name: "Guest",
//             email: "guest@example.com",
//             password: "Admin@4556"
//         }];
//     } catch {
//         return [{id: "user-1", name: "Guest", email: "guest@example.com", password: "password"}];
//     }
// };
//
// const setStoredUsers = (users) => {
//     localStorage.setItem("user", JSON.stringify(users));
// };
//
// export const fetchBooks = async () => {
//     try {
//         await axios.get("http://localhost:4000/api/books").then((res) => {
//             if (res.data && Array.isArray(res.data)) {
//                 books.length = 0;
//                 books.push(...res.data);
//             }
//         });
//     } catch (error) {
//         console.error("Failed to fetch books:", error);
//     }
//     return books;
// };
//
// export const fetchBookById = async (id) => {
//     try {
//         await axios.get(`http://localhost:4000/api/books/${id}`).then((res) => {
//             if (res.data) {
//                 const index = books.findIndex((b) => b.id === id);
//                 if (index !== -1) {
//                     books[index] = res.data;
//                 } else {
//                     books.push(res.data);
//                 }
//             }
//         });
//     } catch (error) {
//         console.error(`Failed to fetch book with id ${id}:`, error);
//     }
//
//     return books.find((b) => b.id === id) || null;
// };
//
// export const login = async ({email, password}) => {
//     try {
//         await axios.post("http://localhost:4000/api/auth/login", {email, password});
//     } catch (error) {
//         console.error("Login failed:", error);
//         throw new Error("Login failed");
//     }
//
//     const users = getStoredUsers();
//     const user = users.find((u) => u.email === email && u.password === password);
//     if (!user) throw new Error("Invalid credentials");
//     localStorage.setItem("auth_user", JSON.stringify({id: user.id, name: user.name, email: user.email}));
//     return {id: user.id, name: user.name, email: user.email};
// };
//
// export const register = async ({name, email, password}) => {
//     try {
//         await axios.post("http://localhost:4000/api/auth/signup", {name, email, password});
//     } catch (error) {
//         console.error("Registration failed:", error);
//         throw new Error("Registration failed");
//     }
//
//     const users = getStoredUsers();
//     if (users.find((u) => u.email === email)) throw new Error("Email already used");
//     const newUser = {id: `user-${Date.now()}`, name, email, password};
//     users.push(newUser);
//     setStoredUsers(users);
//     localStorage.setItem("auth_user", JSON.stringify({id: newUser.id, name: newUser.name, email: newUser.email}));
//     return {id: newUser.id, name: newUser.name, email: newUser.email};
// };
//
// export const logout = () => {
//     localStorage.removeItem("auth_user");
// };
//
// export const getCurrentUser = () => {
//     try {
//         const raw = localStorage.getItem("auth_user");
//         return raw ? JSON.parse(raw) : null;
//     } catch {
//         return null;
//     }
// };
//
// export const getCart = () => {
//     try {
//         const raw = localStorage.getItem("cart");
//         return raw ? JSON.parse(raw) : [];
//     } catch {
//         return [];
//     }
// };
//
// export const updateCart = async (cart) => {const API_URL = "http://localhost:4000";
//
// export const getCurrentUser = async () => {
//     const response = await fetch(`${API_URL}/auth/me`, {
//         credentials: "include",
//     });
//     if (!response.ok) throw new Error("Failed to fetch user");
//     return response.json();
// };
//
// export const login = async ({ email, password }) => {
//     const response = await fetch(`${API_URL}/auth/login`, {
//         method: "POST",
//         headers: { "Content-Type": "application/json" },
//         body: JSON.stringify({ email, password }),
//         credentials: "include",
//     });
//     if (!response.ok) throw new Error("Login failed");
//     return response.json();
// };
//     localStorage.setItem("cart", JSON.stringify(cart));
//     return cart;
// };
//
// export const createOrder = async ({user, items, address}) => {
//     try {
//         await axios.post("http://localhost:4000/api/orders", {user, items, address});
//     } catch (error) {
//         console.error("Order creation failed:", error);
//         throw new Error("Order creation failed");
//     }
//
//     if (!items || items.length === 0) throw new Error("Cart is empty");
//     const orderId = `order-${Date.now()}`;
//     localStorage.removeItem("cart");
//     return {orderId, user, items, address, createdAt: new Date().toISOString()};
// };
//
// export const isAuthenticated = () => {
//     return getCurrentUser() !== null;
// };


import axios from "axios";

const API_URL = "http://localhost:4000/api";

const api = axios.create({
    baseURL: API_URL,
    withCredentials: true,
    headers: {
        "Content-Type": "application/json",
    },
});


export const fetchBooks = async () => {
    const { data } = await api.get("/books");
    return data
};

export const fetchBookById = async (id) => {
    const { data } = await api.get("/book", {
        params: { id },
    });
    return data;
};


export const login = async ({ email, password }) => {
    const { data } = await api.post("/auth/login", { email, password });
    localStorage.setItem("auth_user", JSON.stringify(data));
};

export const register = async ({ name, email, password }) => {
    const { data } = await api.post("/auth/signup", {
        name,
        email,
        password,
    });
    return data;
};

export const logout = async () => {
    localStorage.removeItem("auth_user");
};

export const getCurrentUser = () => {
    try {
        const raw = localStorage.getItem("auth_user");
        return raw ? JSON.parse(raw) : null;
    } catch {
        return null;
    }
};

export const isAuthenticated = async () => {
    try {
        await getCurrentUser();
        return true;
    } catch {
        return false;
    }
};


export const getCart = () => {
    try {
        return JSON.parse(localStorage.getItem("cart")) || [];
    } catch {
        return [];
    }
};

export const updateCart = (cart) => {
    localStorage.setItem("cart", JSON.stringify(cart));
    return cart;
};


export const createOrder = async ({ items, address }) => {
    if (!items || items.length === 0) {
        throw new Error("Cart is empty");
    }

    const { data } = await api.post("/orders", {
        items,
        address,
    });

    localStorage.removeItem("cart");
    return data;
};
