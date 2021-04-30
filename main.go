package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var mysqlDB *sql.DB

func Server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/api/products", fetchProducts).Methods("GET")
	return router
}

func main() {
	mysqlDB = connect()
	defer mysqlDB.Close()
	router := Server()
	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  200,
		"message": "Hello, welcome",
	})
}

func fetchProducts(writer http.ResponseWriter, request *http.Request) {
	// row, err := mysqlDB.Query("SELECT * FROM products")
	// if err != nil {
	// 	renderJSON(writer, map[string]interface{}{
	// 		"message": "Not found",
	// 	})
	// }
	// renderJSON(writer, map[string]interface{}{
	// 	"message": "products",
	// })
}
