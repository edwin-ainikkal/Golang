package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : homepage")
	fmt.Fprintf(w, "Hello, this is my first Go web server!")

}

func returnAllproducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : Return all products")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.URL.Path) //to get id
	vars := mux.Vars(r)
	//key := r.URL.Path[len("/product/"):]
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)

		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllproducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/foo", homepage)
	http.ListenAndServe("localhost:10000", myRouter) //normally myrouter was nil
}

func main() {
	Products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00}}
	handleRequests()

}
