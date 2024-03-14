package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hello from Naman")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	router := chi.NewRouter()
 
  router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,  
  }))

 
  v1Router := chi.NewRouter()
  v1Router.Get("/healthz",handlerReadiness)
  
  v1Router.Get("/err", handlerError)

  router.Mount("/v1", v1Router)
  srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}
  
  log.Printf("Server starting on port %v",PORT)
  log.Fatal(srv.ListenAndServe())
}
