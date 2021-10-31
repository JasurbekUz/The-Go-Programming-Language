package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

type Food struct {
	Id    uuid.UUID
	Name  string
	Price float64
}

type Id struct {
	Id uuid.UUID
}

var foods []Food

func main() {

	foods = []Food{}

	http.HandleFunc("/foods", FoodController)

	http.ListenAndServe(":4000", nil)

}

func FoodController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	if r.Method == "GET" { // GET

		encoder.Encode(foods)

	} else if r.Method == "POST" { // POST

		newFood := Food{}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		json.Unmarshal(body, &newFood)

		newFood.Id = uuid.NewV4()

		foods = append(foods, newFood)

		encoder.Encode(newFood)
		fmt.Printf("%T,  %v", newFood.Id, newFood.Id)

	} else if r.Method == "DELETE" { // DELETE

		//deletedFood := Food{}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		id := Id{}

		json.Unmarshal(body, &id)

		deletedFood := Food{}

		foods = func(slc []Food, id uuid.UUID) (nl []Food) {

			nl = []Food{}

			for _, value := range slc {

				if value.Id == id {

					deletedFood = value
					continue
				}

				nl = append(nl, value)
			}

			return nl
		}(foods, id.Id)

		encoder.Encode(deletedFood)
	} else if r.Method == "PUT" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		reFood := Food{}

		json.Unmarshal(body, &reFood)

		foods = func(slc []Food, reFood Food) (nl []Food) {

			nl = []Food{}

			for _, value := range slc {

				if value.Id == reFood.Id {

					value.Id = reFood.Id
					value.Name = reFood.Name
					value.Price = reFood.Price
				}

				nl = append(nl, value)
			}

			return nl
		}(foods, reFood)

		encoder.Encode(reFood)
	}

}
