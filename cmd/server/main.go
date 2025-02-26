package main

import (
	"itmo-devops-fp1/internal/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Server is starting...")

	// Создаем новый роутер
	r := chi.NewRouter()

	// Добавляем middleware (опционально)
	r.Use(middleware.Logger)

	// Регистрируем маршруты
	r.Route("/api/v0", func(r chi.Router) {
		r.Post("/prices", handler.UploadHandler)
		r.Get("/prices", handler.DownloadHandler)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	log.Println("Server started on :8080")
}
