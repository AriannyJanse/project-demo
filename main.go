package main

import (
	"fmt"
	"log"
	"net/http"
	"project-demo/api/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

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

	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	fmt.Println("\nListening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOrigins, corsHeaders, corsMethods)(router)))

}
