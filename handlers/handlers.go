package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	customer "github.com/Tanmay451/mini-project-h2-2023-server/models/customer"
	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	c := customer.NewCustomer("tanmay", time.Date(2000, 11, 14, 10, 45, 16, 0, time.UTC))

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(c)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func GetCustomerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["id"])
}
