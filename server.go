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

// Initiate DB and Error variable
var db *gorm.DB
var err error

// Food Model
type Food struct {
	gorm.Model
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Taste  string `json:"taste"`
}

// Initialize Database with Food Table
func initialMigration() {

	// Check For Connection
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

	// Check For Connection
	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Retrieve Foods, Database Connection Failed")
	}
	defer db.Close()

	// Find All Food Items
	var foods []Food
	db.Find(&foods)

	// Return All Food Items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

// Get specific Food Item
func getFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: getFood")

	// Check For Connection
	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Retrieve Food, Database Connection Failed")
	}
	defer db.Close()

	// Retrieve ID from url variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Find Food Item with the given ID
	var food Food
	db.Where("id = ?", id).Find(&food)

	// Return The Food Item
	json.NewEncoder(w).Encode(food)
}

// Create Food Item
func createFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: createFood")

	// Check For Connection
	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed To Create Food, Database Connection Failed")
	}
	defer db.Close()

	// Decode the raw Data from response Body into food object
	var food Food
	_ = json.NewDecoder(r.Body).Decode(&food)

	// Create Item
	db.Create(&Food{Name: food.Name, Origin: food.Origin, Taste: food.Taste})

	fmt.Fprintf(w, "Sucessfuly Created Food Item")
}

// Delete Food Item
func deleteFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: deleteFood")

	// Check For Connection
	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed To Delete, Database Connection Failed")
	}
	defer db.Close()

	// Retrieve ID from url variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Find Food Item with the Given ID
	var food Food
	db.Where("id = ?", id).Find(&food)

	// Delete the food Item
	db.Delete(&food)

	fmt.Fprintf(w, "Sucessfuly Deleted Food with ID :"+id)
}

// Update Food Item
func updateFood(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: updateFood")

	// Check For Connection
	db, err = gorm.Open("sqlite3", "food.db")
	if err != nil {
		panic("Failed to Update Food, Database Connection Failed")
	}
	defer db.Close()

	// Retrieve ID from url variables
	vars := mux.Vars(r)
	id := vars["id"]

	// Decode new Updated Item into Temporary food Object
	var Tempfood Food
	_ = json.NewDecoder(r.Body).Decode(&Tempfood)

	// Find the food item to be Updated Using ID
	var food Food
	db.Where("id = ?", id).Find(&food)

	// Apply changes to the Food Item
	food.Name = Tempfood.Name
	food.Origin = Tempfood.Origin
	food.Taste = Tempfood.Taste

	// Save it into Database
	db.Save(&food)

	fmt.Fprintf(w, "Sucessfuly Updated Food Item with ID: "+id)
}

// Main Function
func main() {

	// Initialize Database
	initialMigration()

	// Router
	router := mux.NewRouter()

	// Router Handlers
	router.HandleFunc("/api/food", getFoods).Methods("GET")
	router.HandleFunc("/api/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/food/", createFood).Methods("POST")
	router.HandleFunc("/api/food/{id}/", updateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", deleteFood).Methods("DELETE")

	// Listen
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
