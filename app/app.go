package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	mux := mux.NewRouter()
	//define a route
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//start the server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
