package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

type Product struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var products []Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
}

func main() {
	products = append(products, Product{ID: "1", Title: "Apple"})
	products = append(products, Product{ID: "2", Title: "Orange"})
	products = append(products, Product{ID: "3", Title: "Corn"})

	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
