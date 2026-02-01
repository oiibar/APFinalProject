export const books = [
    {
        id: 1,
        title: "Clean Code",
        author: "Robert C. Martin",
        price: 29.99,
        image: "https://images.unsplash.com/photo-1524995997946-a1c2e315a42f"
    },
    {
        id: 2,
        title: "The Pragmatic Programmer",
        author: "Andrew Hunt",
        price: 34.99,
        image: "https://images.unsplash.com/photo-1512820790803-83ca734da794"
    },
    {
        id: 3,
        title: "Introduction to Algorithms",
        author: "CLRS",
        price: 49.99,
        image: "https://images.unsplash.com/photo-1519681393784-d120267933ba"
    }
];

export const fakeLogin = (email, password) => {
    if (email && password) {
        localStorage.setItem("auth", "true");
        return true;
    }
    return false;
};

export const isAuthenticated = () => {
    return localStorage.getItem("auth") === "true";
};

export const logout = () => {
    localStorage.removeItem("auth");
};