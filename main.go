package main

import (
	"log"
	"net/http"

	"github.com/arkiralor/GoURLShortener/pkg/routers"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	godotenv.Load(".env")
	router := mux.NewRouter()
	routers.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:7000", router))
}
