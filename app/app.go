package app

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/Tanmay451/mini-project-h2-2023-server/handlers"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	router.HandleFunc("/getAllCustomer", handlers.GetAllCustomer)
	log.Fatal(http.ListenAndServe(":8000", router))
}
