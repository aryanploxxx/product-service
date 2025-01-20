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

func GetProducts(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("GET Method Called")

		rows, err := database.Query("SELECT * FROM product")
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer rows.Close()

		allProducts := []models.Product{}

		for rows.Next() {
			var oneProduct models.Product
			err := rows.Scan(&oneProduct.Product_id, &oneProduct.Product_name, &oneProduct.Product_quantity)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			allProducts = append(allProducts, oneProduct)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(allProducts)

		json.NewEncoder(w).Encode(allProducts)

	}
}

func GetProductsByID(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET Method By ID Called")

		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("calling for id ", id)

		var oneProduct models.Product
		err := database.QueryRow("Select * FROM product WHERE Product_id = $1", id).Scan(&oneProduct.Product_id, &oneProduct.Product_name, &oneProduct.Product_quantity)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		jsonData, err := json.Marshal(oneProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}

func CreateProduct(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST Method Called")

		var newProduct models.Product

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &newProduct)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = database.QueryRow("INSERT INTO product (product_name, product_quantity) VALUES ($1, $2) RETURNING Product_id", newProduct.Product_name, newProduct.Product_quantity).Scan(&newProduct.Product_id)
		if err != nil {
			http.Error(w, "Error saving product", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(newProduct)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	}
}
