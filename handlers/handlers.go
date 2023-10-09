package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	c := []Customer{
		{
			Name: "tanmay",
			Age:  25,
		},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(c)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}