package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full name"  xml:"name"'`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

	//encode the customers to xml

}
