package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Food struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

// Foods Data
var Foods = []Food{

	Food{ID: 1, Name: "Biryani", Origin: "Persia"},
	Food{ID: 2, Name: "Biryani", Origin: "Persia"},
	Food{ID: 3, Name: "Biryani", Origin: "Persia"},
	Food{ID: 4, Name: "Biryani", Origin: "Persia"},
	Food{ID: 5, Name: "Biryani", Origin: "Persia"},
	Food{ID: 6, Name: "Biryani", Origin: "Persia"},
	Food{ID: 7, Name: "Biryani", Origin: "Persia"},
}

func getFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: getFoods")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Foods)
}

func main() {

	// Router
	router := mux.NewRouter()

	// Router Handlers
	router.HandleFunc("/api/food", getFoods).Methods("GET")

	// Listen
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
