package main

import (
	"auth/security/authenticator"
	"auth/security/login"
	"auth/security/register"
	"net/http"

	"log"

	"github.com/gorilla/mux"

	"auth/logic"
)

// TO DO: server address debe ser tomada de la variable de entorno segun donde deba ser
const serverAddress = "0.0.0.0:3000"

func main() {
	router := mux.NewRouter()

	// Unprotected routes: login and register
	router.HandleFunc("/api/user/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/api/user/register", register.RegisterHandler).Methods("POST")

	// Protected routes: routes that require authorization: authorization token retrieved in login handler
	router.HandleFunc("/api/authenticate", authenticator.Authenticate).Methods("GET") //ejemplo de funcion que protege paths con auth en paquete pathprotector
	router.HandleFunc("/api/logic/", logic.DoSomeLogic).Methods("GET")

	log.Println("Starting the server on: " + serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatal("Could not start the server", err)
	}
}
