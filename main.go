package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	http.HandleFunc("/getAllCustomer", func(w http.ResponseWriter, r *http.Request) {
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
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
