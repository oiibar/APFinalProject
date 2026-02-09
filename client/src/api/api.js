// Placeholder hard-coded dataset and fake API functions for frontend prototyping
export const books = [
    {
        id: "book-1",
        title: "Введение в Go",
        author: "Иван Иванов",
        price: 19.99,
        image: "https://images.unsplash.com/photo-1512820790803-83ca734da794",
        description: "Практическое введение в Go.",
        pages: 320,
        publisher: "TechPress",
        isbn: "978-1-1111-1111",
        stock: 8,
    },
    {
        id: "book-2",
        title: "React для начинающих",
        author: "Анна Петрова",
        price: 24.5,
        image: "https://images.unsplash.com/photo-1524995997946-a1c2e315a42f",
        description: "Шаг за шагом к SPA.",
        pages: 280,
        publisher: "WebBooks",
        isbn: "978-2-2222-2222",
        stock: 5,
    },
    {
        id: "book-3",
        title: "Алгоритмы и структуры данных",
        author: "К. Смит",
        price: 29.0,
        image: "https://images.unsplash.com/photo-1519681393784-d120267933ba",
        description: "Классика по алгоритмам.",
        pages: 600,
        publisher: "AlgoPub",
        isbn: "978-3-3333-3333",
        stock: 3,
    },
    {
        id: "book-4",
        title: "CSS современные практики",
        author: "Мария Ли",
        price: 18.75,
        image: "https://images.unsplash.com/photo-1493119508027-2b584f234d6c",
        description: "Советы по современному CSS.",
        pages: 200,
        publisher: "DesignHub",
        isbn: "978-4-4444-4444",
        stock: 10,
    },
    {
        id: "book-6",
        title: "Тестирование веб‑приложений",
        author: "Екатерина Новикова",
        price: 16.99,
        image: "https://images.unsplash.com/photo-1496181133206-80ce9b88a853",
        description: "Подходы к тестированию фронтенда.",
        pages: 240,
        publisher: "QApress",
        isbn: "978-6-6666-6666",
        stock: 6,
    },
];

// Simple in-memory users list for demo (will persist to localStorage on register)
const LOCAL_USERS_KEY = "demo_users_v1";
const LOCAL_CART_KEY = "demo_cart_v1";

const getStoredUsers = () => {
    try {
        const raw = localStorage.getItem(LOCAL_USERS_KEY);
        return raw ? JSON.parse(raw) : [{ id: "user-1", name: "Guest", email: "guest@example.com", password: "password" }];
    } catch {
        return [{ id: "user-1", name: "Guest", email: "guest@example.com", password: "password" }];
    }
};

const setStoredUsers = (users) => {
    localStorage.setItem(LOCAL_USERS_KEY, JSON.stringify(users));
};

// API emulation: small delays to simulate network
const delay = (ms = 300) => new Promise((res) => setTimeout(res, ms));

export const fetchBooks = async () => {
    await delay(200);
    return books;
};

export const fetchBookById = async (id) => {
    await delay(150);
    return books.find((b) => b.id === id) || null;
};

export const login = async ({ email, password }) => {
    await delay(200);
    const users = getStoredUsers();
    const user = users.find((u) => u.email === email && u.password === password);
    if (!user) throw new Error("Invalid credentials");
    localStorage.setItem("auth_user", JSON.stringify({ id: user.id, name: user.name, email: user.email }));
    return { id: user.id, name: user.name, email: user.email };
};

export const register = async ({ name, email, password }) => {
    await delay(250);
    const users = getStoredUsers();
    if (users.find((u) => u.email === email)) throw new Error("Email already used");
    const newUser = { id: `user-${Date.now()}`, name, email, password };
    users.push(newUser);
    setStoredUsers(users);
    localStorage.setItem("auth_user", JSON.stringify({ id: newUser.id, name: newUser.name, email: newUser.email }));
    return { id: newUser.id, name: newUser.name, email: newUser.email };
};

export const logout = async () => {
    await delay(50);
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

export const getCart = async () => {
    await delay(100);
    try {
        const raw = localStorage.getItem(LOCAL_CART_KEY);
        return raw ? JSON.parse(raw) : [];
    } catch {
        return [];
    }
};

export const updateCart = async (cart) => {
    await delay(120);
    localStorage.setItem(LOCAL_CART_KEY, JSON.stringify(cart));
    return cart;
};

export const createOrder = async ({ user, items, address }) => {
    await delay(300);
    // simple validation
    if (!items || items.length === 0) throw new Error("Cart is empty");
    const orderId = `order-${Date.now()}`;
    // For demo, just clear cart
    localStorage.removeItem(LOCAL_CART_KEY);
    return { orderId, user, items, address, createdAt: new Date().toISOString() };
};

export const isAuthenticated = () => {
    return getCurrentUser() !== null;
};