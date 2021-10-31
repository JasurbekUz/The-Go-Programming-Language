package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/numbers", Numbers)
	http.HandleFunc("/countries", Countries)
	http.HandleFunc("/companies", Companies)

	http.ListenAndServe(":4000", nil)
}

func Numbers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
}

func Countries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cities := []string{"Uzbekistan", "Kazakhstan", "Tadjikstan", "Kyrgzstan", "Turkmenstan", "Afagnstan"}

	byte, err := json.Marshal(cities)

	if err != nil {

		panic(err)
	}

	fmt.Fprintf(w, string(byte))
}

func Companies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`["Google","Yandex","Apple","Samsung","Amazon","MicroSoft"]`))

	//fmt.Fprintf(w, `["Google","Yandex","Apple","Samsung","Amazon","MicroSoft"]`)

	//[]string{"Google", "Apple", "Amazon", "Yandex"}
}
