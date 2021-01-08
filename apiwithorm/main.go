package main

import (
	"fmt"
	"guilherme/criandowebapigo/apiwithorm/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", controllers.AllUseres).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", controllers.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", controllers.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", controllers.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	fmt.Println("GO ORM Tutorial")
	controllers.InitialMigration()
	handleRequests()
}
