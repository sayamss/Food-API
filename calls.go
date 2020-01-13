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

// Food Model ...
type FoodItem struct {
	ID     int
	Name   string
	Origin string
	Taste  string
}

type Data struct {
	Title string
	items []FoodItem
}

// GetAllFoods ...
func GetAllFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: GetAllFoods")

	switch r.Method {

	case http.MethodPost:

		newItem := FoodItem{
			Name:   r.FormValue("name"),
			Origin: r.FormValue("origin"),
			Taste:  r.FormValue("taste"),
		}

		requestBody, err := json.Marshal(map[string]string{
			"name":   newItem.Name,
			"origin": newItem.Origin,
			"taste":  newItem.Taste,
		})

		if err != nil {
			fmt.Println("Error Occured")
		}

		resp, postErr := http.Post("http://localhost:8000/api/food/", "application/json", bytes.NewBuffer(requestBody))

		if postErr != nil {
			fmt.Println("Error Occured")
		}

		fmt.Println(resp.Body)
	}

	response, err := http.Get("http://localhost:8000/api/food")
	if err != nil {
		fmt.Printf("Could Not Fetch Foods, Error: %s", err)
		return
	}

	defer response.Body.Close()

	var items []FoodItem
	_ = json.NewDecoder(response.Body).Decode(&items)

	templ := template.Must(template.ParseFiles("templates/home.html"))

	templ.Execute(w, items)
	return
}

// CreateNewFood ...
func updateFoodItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: Updating Food Item")
	id := r.FormValue("id")
	updateItem := FoodItem{
		Name:   r.FormValue("name"),
		Origin: r.FormValue("origin"),
		Taste:  r.FormValue("taste"),
	}

	requestBody, err := json.Marshal(map[string]string{
		"name":   updateItem.Name,
		"origin": updateItem.Origin,
		"taste":  updateItem.Taste,
	})

	if err != nil {
		fmt.Println("Error Occured")
	}

	url := "http://localhost:8000/api/food/" + id + "/"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println("error occured")
	}
	fmt.Println(resp.Body)
	http.Redirect(w, r, "/", 301)

}

func main() {

	http.HandleFunc("/", GetAllFoods)
	http.HandleFunc("/update", updateFoodItem)
	http.ListenAndServe(":8080", nil)
}
