package main

import (
	"fmt"
	"log"
	"net/http"
	"project-demo/api/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "The element was not found"}`))
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", controllers.GetAllUsers).Methods(http.MethodGet)
	api.HandleFunc("/users/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	api.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	api.HandleFunc("/users/{id}", controllers.UpdateUserByID).Methods(http.MethodPut)
	api.HandleFunc("/users/{id}", controllers.DeleteUserByID).Methods(http.MethodDelete)

	api.HandleFunc("/companies", controllers.GetAllCompanies).Methods(http.MethodGet)
	api.HandleFunc("/companies/{id}", controllers.GetCompanyByID).Methods(http.MethodGet)
	api.HandleFunc("/companies", controllers.CreateCompany).Methods(http.MethodPost)
	api.HandleFunc("/companies/{id}", controllers.UpdateCompanyByID).Methods(http.MethodPut)
	api.HandleFunc("/companies/{id}", controllers.DeleteCompanyByID).Methods(http.MethodDelete)

	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("\nListening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOrigins, corsHeaders, corsMethods)(router)))

}
