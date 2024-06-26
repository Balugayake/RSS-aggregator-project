package main

import (
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerRediness)

	v1Router.Get("/error", hanlerError)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Printf("Listening on port %v\n", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
