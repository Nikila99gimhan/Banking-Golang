package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {

	//define a route

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//start the server
	log.Fatal(http.ListenAndServe(":8000", nil))

}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")

}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		Customer{
			Name:    "John",
			City:    "New York",
			Zipcode: "10001",
		},
		Customer{
			Name:    "Jane",
			City:    "Colombo",
			Zipcode: "11500",
		},
		Customer{
			Name:    "Joe",
			City:    "New York",
			Zipcode: "10001",
		},
	}
	json.NewEncoder(w).Encode(customers)

}