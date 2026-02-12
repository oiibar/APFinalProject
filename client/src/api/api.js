export const books = [
    {
        id: "book-1",
        title: "Book 1",
        author: "John Doe",
        price: 9.99,
        image: "https://placehold.co/70x70",
        description: "Book 1 description.",
        pages: 200,
        publisher: "Publisher 1",
        stock: 8,
    },
    {
        id: "book-2",
        title: "Book 2",
        author: "John Doe",
        price: 9.99,
        image: "https://placehold.co/70x70",
        description: "Book 2 description.",
        pages: 300,
        publisher: "Publisher 2",
        stock: 5,
    },
    {
        id: "book-3",
        title: "Book 3",
        author: "John Doe",
        price: 9.99,
        image: "https://placehold.co/70x70",
        description: "Book 3 description.",
        pages: 400,
        publisher: "Publisher 3",
        stock: 3,
    },
    {
        id: "book-4",
        title: "Book 4",
        author: "John Doe",
        price: 9.99,
        image: "https://placehold.co/70x70",
        description: "Book 4 description.",
        pages: 500,
        publisher: "Publisher 4",
        stock: 10,
    },
    {
        id: "book-6",
        title: "Book 5",
        author: "John Doe",
        price: 9.99,
        image: "https://placehold.co/70x70",
        description: "Book 5 description.",
        pages: 100,
        publisher: "Publisher 5",
        stock: 6,
    },
];

const getStoredUsers = () => {
    try {
        const raw = localStorage.getItem("user");
        return raw ? JSON.parse(raw) : [{
            id: "user-1",
            name: "Guest",
            email: "guest@example.com",
            password: "password"
        }];
    } catch {
        return [{id: "user-1", name: "Guest", email: "guest@example.com", password: "password"}];
    }
};

const setStoredUsers = (users) => {
    localStorage.setItem("user", JSON.stringify(users));
};

export const fetchBooks = () => {
    return books;
};

export const fetchBookById = (id) => {
    return books.find((b) => b.id === id) || null;
};

export const login = ({email, password}) => {
    const users = getStoredUsers();
    const user = users.find((u) => u.email === email && u.password === password);
    if (!user) throw new Error("Invalid credentials");
    localStorage.setItem("auth_user", JSON.stringify({id: user.id, name: user.name, email: user.email}));
    return {id: user.id, name: user.name, email: user.email};
};

export const register = async ({name, email, password}) => {
    const users = getStoredUsers();
    if (users.find((u) => u.email === email)) throw new Error("Email already used");
    const newUser = {id: `user-${Date.now()}`, name, email, password};
    users.push(newUser);
    setStoredUsers(users);
    localStorage.setItem("auth_user", JSON.stringify({id: newUser.id, name: newUser.name, email: newUser.email}));
    return {id: newUser.id, name: newUser.name, email: newUser.email};
};

export const logout = () => {
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

export const getCart = () => {
    try {
        const raw = localStorage.getItem("cart");
        return raw ? JSON.parse(raw) : [];
    } catch {
        return [];
    }
};

export const updateCart = (cart) => {
    localStorage.setItem("cart", JSON.stringify(cart));
    return cart;
};

export const createOrder = ({user, items, address}) => {
    if (!items || items.length === 0) throw new Error("Cart is empty");
    const orderId = `order-${Date.now()}`;
    localStorage.removeItem("cart");
    return {orderId, user, items, address, createdAt: new Date().toISOString()};
};

export const isAuthenticated = () => {
    return getCurrentUser() !== null;
};