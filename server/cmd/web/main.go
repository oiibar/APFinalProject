package main

import (
	"database/sql"
	"final/internal/handlers"
	"final/internal/middleware"
	"final/internal/routes"
	"final/internal/store"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	log_channel chan string
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:     infoLog,
		errorLog:    errorLog,
		log_channel: make(chan string, 256),
	}
	go app.log_worker()

	mux := http.NewServeMux()
	q := make(chan int, 100)

	// Load JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		errorLog.Fatal("JWT_SECRET is required (set it in your environment)")
	}
	handlers.JwtKey = []byte(jwtSecret)
	middleware.JwtKey = []byte(jwtSecret)

	// Load database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		errorLog.Fatal("DATABASE_URL is required (set it in your environment)")
	}

	// Connect to PostgreSQL
	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Create tables if they do not exist
	if err := createTables(db); err != nil {
		errorLog.Fatal(err)
	}

	// Use PostgreSQL store (data will be persisted)
	st := store.NewPostgresStore(db)

	// Ensure there is at least one admin user for testing/demo
	store.EnsureDefaultAdmin(st)

	// Background worker: simulates async order processing and updates status
	go func() {
		for id := range q {
			time.Sleep(750 * time.Millisecond)
			_ = st.UpdateOrderStatus(id, "completed")
		}
	}()

	handler := routes.Routes(mux, st, q)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  handler,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func (app *application) log_worker() {
	for msg := range app.log_channel {
		app.infoLog.Println(msg)
	}
}

// openDB establishes a PostgreSQL connection
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// createTables creates required tables if they do not exist
func createTables(db *sql.DB) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user',
			created_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);`,

		`CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			price NUMERIC(10,2) NOT NULL CHECK (price > 0),
			image TEXT,
			description TEXT,
			category TEXT,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);`,

		`CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			status TEXT NOT NULL DEFAULT 'pending',
			total NUMERIC(10,2) NOT NULL DEFAULT 0,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);`,

		`CREATE TABLE IF NOT EXISTS order_items (
			id SERIAL PRIMARY KEY,
			order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
			book_id INT NOT NULL REFERENCES books(id),
			quantity INT NOT NULL CHECK (quantity > 0),
			unit_price NUMERIC(10,2) NOT NULL CHECK (unit_price > 0)
		);`,
	}

	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	return nil
}
