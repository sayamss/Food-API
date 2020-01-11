package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Food Model
type Food struct {
	gorm.Model
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Taste  string `json:"taste"`
}

func printAllFoods() {

	res, err := http.Get("https://localhost:8000/api/food")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}

func getSpecificFood(id string) {

	url := "https://localhost:8000/api/food/" + id

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}

// Create Food
func CreateFood(name string, origin string, taste string) {

	var food Food
	food.Name = name
	food.Origin = origin
	food.Taste = taste

}

func main() {
	printAllFoods()
	getSpecificFood("1")
}
