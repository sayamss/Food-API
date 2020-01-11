package main

// Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// Food Model
type Food struct {
	gorm.Model
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Taste  string `json:"taste"`
}

func initialMigration() {

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection to Database Failed")
	}

	defer db.Close()

	db.AutoMigrate(&Food{})
}

// Get All Foods
func getFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: getFoods")

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Retrieve Foods, Database Connection Failed")
	}
	defer db.Close()

	var foods []Food
	db.Find(&foods)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: getFood")

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Retrieve Food, Database Connection Failed")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var food Food
	db.Where("id = ?", id).Find(&food)

	json.NewEncoder(w).Encode(food)
}

func createFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: createFood")

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed To Create Food, Database Connection Failed")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	origin := vars["origin"]
	taste := vars["taste"]

	db.Create(&Food{Name: name, Origin: origin, Taste: taste})
	fmt.Fprintf(w, "Successfuly Created Food Item")
}

func deleteFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: deleteFood")

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed To Delete, Database Connection Failed")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var food Food
	db.Where("id = ?", id).Find(&food)
	db.Delete(&food)

	fmt.Fprintf(w, "Sucessfuly Deleted Food with ID :"+id)
}

func updateFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: updateFood")

	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed to Update Food, Database Connection Failed")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	name := vars["name"]
	origin := vars["origin"]
	taste := vars["taste"]

	var food Food
	db.Where("id = ?", id).Find(&food)

	food.Name = name
	food.Origin = origin
	food.Taste = taste

	db.Save(&food)

	fmt.Fprintf(w, "Sucessfuly Updated Food Item with ID: "+id)
}

// Main Function
func main() {

	initialMigration()
	// Router
	router := mux.NewRouter()

	// Router Handlers
	router.HandleFunc("/api/food", getFoods).Methods("GET")
	router.HandleFunc("/api/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/food/{name}/{origin}/{taste}", createFood).Methods("POST")
	router.HandleFunc("/api/food/{id}/{name}/{origin}/{taste}", updateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", deleteFood).Methods("DELETE")

	// Listen
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
