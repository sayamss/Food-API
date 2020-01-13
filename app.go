// calls.go
package main

// Imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// FoodItem ...
type FoodItem struct {
	ID     int
	Name   string
	Origin string
	Taste  string
}

// GetAllFoods ...
func GetAllFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: GetAllFoods")

	// Post Request
	if r.Method == http.MethodPost {

		// Get the values from Form Data
		requestBody, err := json.Marshal(map[string]string{
			"name":   r.FormValue("name"),
			"origin": r.FormValue("origin"),
			"taste":  r.FormValue("taste"),
		})

		// Error Check
		if err != nil {
			fmt.Println("Error Occured")
			return
		}

		// Send the post request with the required Values
		resp, postErr := http.Post("http://localhost:8000/api/food/", "application/json", bytes.NewBuffer(requestBody))

		if postErr != nil {
			fmt.Println("Error Occured")
			return
		}

		fmt.Println(resp.Body)
	}

	// Fetch All the Food Items
	response, err := http.Get("http://localhost:8000/api/food")
	if err != nil {
		fmt.Printf("Could Not Fetch Foods, Error: %s", err)
		return
	}

	defer response.Body.Close()

	// Store the Fetched Food Items in an Array
	var items []FoodItem
	_ = json.NewDecoder(response.Body).Decode(&items)

	// Template
	templ := template.Must(template.ParseFiles("templates/home.html"))

	templ.Execute(w, items)
	return
}

// updateFoodItem ...
func updateFoodItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Updating Food Item")

	// Retrieve Food Item ID
	id := r.FormValue("id")

	// New Values
	requestBody, err := json.Marshal(map[string]string{
		"name":   r.FormValue("name"),
		"origin": r.FormValue("origin"),
		"taste":  r.FormValue("taste"),
	})

	// Error Check
	if err != nil {
		fmt.Println("Error Occured")
		return
	}

	// Specific URL for updating the required Item
	url := "http://localhost:8000/api/food/" + id + "/"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println("error occured")
		return
	}
	fmt.Println(resp.Body)

	// Redirect to Home Page
	http.Redirect(w, r, "/", 301)

}

func deleteFoodItem(w http.ResponseWriter, r *http.Request) {

	// Send Post request to the URL with id to delete the Item
	id := r.FormValue("id")

	url := "http://localhost:8000/api/food/delete/" + id

	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		fmt.Println("error occured")
		return
	}

	fmt.Println(resp.Body)

	// Redirect to Home Page
	http.Redirect(w, r, "/", 301)
}

func main() {

	// Handler functions
	http.HandleFunc("/", GetAllFoods)
	http.HandleFunc("/update", updateFoodItem)
	http.HandleFunc("/delete", deleteFoodItem)
	http.ListenAndServe(":8080", nil)
}
