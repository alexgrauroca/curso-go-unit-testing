package main

import (
	"catching-pokemons/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", controller.GetPokemon).Methods("GET")

	log.Println("Server listening on port :8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Error found")
	}
}
