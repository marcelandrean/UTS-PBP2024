package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	// "github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	// router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users/login", controllers.Login).Methods("POST")

	// CORS Handling
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:8888"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", handler))
}
