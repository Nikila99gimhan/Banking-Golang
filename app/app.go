package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()
	//define a route

	router.HandleFunc("/greet", greet).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")

	//start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST Request Received")
}
