# Food RestAPI Using Go
****************************************************
## Tech Stack
* [GoLang](https://golang.org/)
* [mux](https://github.com/gorilla/mux)
* [gorm](https://github.com/jinzhu/gorm)
* [sqlite](https://www.sqlite.org/index.html)
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
  2. Method **POST**

#### Delete Food Item by ID
* **/api/food/{id}** -
  1. Deletes Food Item with the specified ID
  2. Method - **DELETE**

#### Update Food Item by ID
* **/api/food/{id}** -
  1. JSON Example: { "name": "IceCream 2", "origin":"Belgium", "taste": "10"}, url - /api/food/1
  2. Method - **PUT**
****************************************************
  
  
