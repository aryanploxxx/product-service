package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"product-service/handler"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var database *sql.DB

func init() {
	dbURL := "host=localhost port=5432 user=postgres password=Pyari@123 sslmode=disable dbname=productDB"
	fmt.Println("Database URL: ", dbURL)

	var err error
	database, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Println("Database connected successfully")
	}

	_, _ = database.Exec("CREATE TABLE IF NOT EXISTS product (product_id SERIAL PRIMARY KEY, product_name TEXT, product_quantity TEXT)")
	_, _ = database.Exec("CREATE TABLE IF NOT EXISTS customer (customer_id SERIAL PRIMARY KEY, customer_name TEXT, customer_email TEXT)")
}

func main() {

	defer database.Close()

	router := mux.NewRouter()

	router.HandleFunc("/products", handler.GetProducts(database)).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetProductsByID(database)).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct(database)).Methods("POST")

	router.HandleFunc("/customers", handler.GetCustomers(database)).Methods("GET")
	router.HandleFunc("/customers/{id}", handler.GetCustomersByID(database)).Methods("GET")
	router.HandleFunc("/customers", handler.CreateCustomer(database)).Methods("POST")

	http.ListenAndServe(":8080", router)

}
