package main

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Food Model
type Food struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

// Foods Data
var Foods = []Food{
	Food{ID: "1", Name: "Biryani", Origin: "Persia"},
	Food{ID: "2", Name: "Biryani", Origin: "Persia"},
	Food{ID: "3", Name: "Biryani", Origin: "Persia"},
	Food{ID: "4", Name: "Biryani", Origin: "Persia"},
	Food{ID: "5", Name: "Biryani", Origin: "Persia"},
	Food{ID: "6", Name: "Biryani", Origin: "Persia"},
	Food{ID: "7", Name: "Biryani", Origin: "Persia"},
}

// Get All Foods
func getFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: getFoods")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	pars := mux.Vars(r)

	// Find Food
	for _, f := range Foods {
		if pars["id"] == f.ID {
			json.NewEncoder(w).Encode(f)
			return
		}
	}
}

func createFood(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var food Food
	_ = json.NewDecoder(r.Body).Decode(&food)

	food.ID = strconv.Itoa(rand.Intn(1000))
	Foods = append(Foods, food)

	json.NewEncoder(w).Encode(food)
}

func deleteFood(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	pars := mux.Vars(r)

	for i, f := range Foods {

		if pars["id"] == f.ID {
			Foods = append(Foods[:i], Foods[i+1:]...)
			break
		}
	}
}

func updateFood(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	pars := mux.Vars(r)

	var food Food
	_ = json.NewDecoder(r.Body).Decode(&food)

	for i, f := range Foods {

		if pars["id"] == f.ID {

			var tempID string = Foods[i].ID
			Foods[i] = food
			Foods[i].ID = tempID

		}
	}
	json.NewEncoder(w).Encode(Foods)
}

// Main Function
func main() {

	// Router
	router := mux.NewRouter()

	// Router Handlers
	router.HandleFunc("/api/food", getFoods).Methods("GET")
	router.HandleFunc("/api/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/food/", createFood).Methods("POST")
	router.HandleFunc("/api/food/{id}", updateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", deleteFood).Methods("DELETE")

	// Listen
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
