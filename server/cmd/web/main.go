package main

import (
	"final/internal/handlers"
	"final/internal/middleware"
	"final/internal/routes"
	"final/internal/store"
	"flag"
	"log"
	"net/http"
	"os"
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
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	go app.log_worker()

	mux := http.NewServeMux()
	q := make(chan int, 100)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secretus"
	}
	handlers.JwtKey = []byte(jwtSecret)
	middleware.JwtKey = []byte(jwtSecret)

	var st store.Store
	st = store.NewMemoryStore()

	handler := routes.Routes(mux, st, q)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  handler,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func (app *application) log_worker() {
	for msg := range app.log_channel {
		app.infoLog.Println(msg)
	}
}
