package main

import (
    "net/http"
    "os"
    "log"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/joho/godotenv"
    "github.com/aidosgal/nextgen/handler"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

    port := os.Getenv("APP_PORT")

    router := chi.NewMux()

    router.Use(middleware.Logger)

    router.Handle("/*", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))

    router.Get("/", handler.Make(handler.HandleHomeShow))

    log.Println("Server started on port", port)

    err = http.ListenAndServe(port, router)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
