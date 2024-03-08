package main

import (
	"fmt"
	"log"
	"net/http"
	"uts/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	// "github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/rooms/detail/{id}", controllers.GetDetailRoom).Methods("GET")
	router.HandleFunc("/participants", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/participants/{id}", controllers.LeaveRoom).Methods("DELETE")

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
