package app

import (
	"fmt"

	"net/http"

	handlers "github.com/Tanmay451/mini-project-h2-2023-server/handlers"

	log "github.com/Tanmay451/mini-project-h2-2023-server/modules/log"
	logrus "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	}).
		Methods(http.MethodGet).
		Schemes("http", "https")

	customerRouter := router.PathPrefix("/customer").Subrouter()

	customerRouter.HandleFunc("/", handlers.GetCustomer).
		Methods(http.MethodGet).
		Schemes("http", "https")

	customerRouter.HandleFunc("/{id:[0-9]+}", handlers.GetCustomer).
		Methods(http.MethodGet).
		Schemes("http", "https")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Logger.WithFields(logrus.Fields{
			"status": "stop",
			"error":  err,
		}).Panicln("App is terminated")
	}
}
