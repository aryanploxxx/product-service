package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"product-service/models"

	"github.com/gorilla/mux"
)

func GetCustomers(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("GET Method Called")

		rows, err := database.Query("SELECT * FROM customer")
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer rows.Close()

		allCustomers := []models.Customer{}

		for rows.Next() {
			var oneCustomer models.Customer
			err := rows.Scan(&oneCustomer.Customer_id, &oneCustomer.Customer_name, &oneCustomer.Customer_email)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			allCustomers = append(allCustomers, oneCustomer)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(allCustomers)

		json.NewEncoder(w).Encode(allCustomers)

	}
}

func GetCustomersByID(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET Method By ID Called")

		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("calling for id ", id)

		var oneCustomer models.Customer
		err := database.QueryRow("Select * FROM customer WHERE Customer_id = $1", id).Scan(&oneCustomer.Customer_id, &oneCustomer.Customer_name, &oneCustomer.Customer_email)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		jsonData, err := json.Marshal(oneCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}

func CreateCustomer(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST Method Called")

		var newCustomer models.Customer

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &newCustomer)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = database.QueryRow("INSERT INTO customer (customer_name, customer_email) VALUES ($1, $2) RETURNING Customer_id", newCustomer.Customer_name, newCustomer.Customer_email).Scan(&newCustomer.Customer_id)
		if err != nil {
			http.Error(w, "Error saving customer", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(newCustomer)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	}
}
