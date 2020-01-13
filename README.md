# Food RestAPI Using Go
****************************************************
## Tech Stack
* [GoLang](https://golang.org/)
* [mux](https://github.com/gorilla/mux)
* [gorm](https://github.com/jinzhu/gorm)
* [sqlite](https://www.sqlite.org/index.html)

## How To Run

* Install Go,mux,gorm( you can find the installation procedure in the links above )
* Run **go run server.go**
* Run **go run app.go**
* Go to **localhost:8080/home**
* You can choose to delete the food.db and start fresh with an empty database

## UI application On **localhost:8080/home** 
![ScreenShot](https://github.com/sayamss/Food-API/blob/master/screenshot/Screenshot%20from%202020-01-13%2022-12-48.png)

## RestAPI Endpoints
****************************************************
### Root Url - localhost:8000

****************************************************
## GET requests

#### Get All Food Items
* **/api/food** - 
  1. Returns all the food items in the database
  2. Method - **GET**

#### Get Food Item By ID
* **/api/food/{id}** -
  1. Returns Food Item with the specified ID
  2. Method - **GET**
  
****************************************************
## POST requests

#### Create Food Item
* **/api/food/** -
  1. JSON Example: { 	"name": "Tea","origin": "China","taste": "10"}
  2. Method - **POST**

#### Delete Food Item by ID
* **/api/food/delete/{id}** -
  1. Deletes Food Item with the specified ID
  2. Method - **POST**

#### Update Food Item by ID
* **/api/food/{id}/** -
  1. JSON Example: { "name": "IceCream 2", "origin":"Belgium", "taste": "10"}, url - /api/food/1
  2. Method - **POST**
****************************************************
  
  
