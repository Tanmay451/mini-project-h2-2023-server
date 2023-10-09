package app

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/Tanmay451/mini-project-h2-2023-server/handlers"
)

func Start() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/getAllCustomer", handlers.GetAllCustomer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
