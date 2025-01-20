# Product and Consumer Management Service

## Directory Structure 
```
product-service/
├── go.mod
├── go.sum
├── handler/
│   ├── allCustomerHandlers.go
│   └── allProductHandlers.go
├── main.go
└── models/
    ├── customer.go
    └── product.go
```

## Project Setup

1. Setup Dependencies
``` golang
  go get "github.com/lib/pq" "github.com/gorilla/mux"
```
2. Run Main Program
``` golang
  go run main.go
```
4. Make Database: ``` productDB ``` in PostgreSQL

## API Endpoints

### Get All Products
- Endpoint: /products
- Method: GET
- Description: Creates a new product.
  
![image](https://github.com/user-attachments/assets/5fb965c8-8531-4a17-aef9-8bbbc7bf7c83)

### Get Product by ID
- Endpoint: /products/{id}
- Method: GET
- Description: Fetch Product by ID.

![image](https://github.com/user-attachments/assets/5ef0eead-37d7-49b9-a24a-1e440c3c5776)

### Create New Product
- Endpoint: /products
- Method: POST
- Description: Creates a new product.
```sh
    {
      "Product_name": "Notebook",
      "Product_quantity": 100
    }
 ```

![image](https://github.com/user-attachments/assets/01fce64a-414e-4bd7-a529-5a6ba31e704c)

### Get All Customers
- Endpoint: /customers
- Method: GET
- Description: Get all customers.
  
![image](https://github.com/user-attachments/assets/51a95353-4c33-4f46-8d98-427ba494bc79)

### Get Customers by ID
- Endpoint: /customers/{id}
- Method: GET
- Description: Fetch Customer by ID.

![image](https://github.com/user-attachments/assets/9338196e-d1d9-48a4-84aa-bafc6390f479)

### Create New Customer
- Endpoint: /customers
- Method: POST
- Description: Creates a new customer.
```sh
    {
      "Customer_name": "Aryaman",
      "Customer_email": "aryaman@gmail.com"
    }
 ```

![image](https://github.com/user-attachments/assets/c0ad2167-c12a-4ae7-af3f-58c4cf95486d)
