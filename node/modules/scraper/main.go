package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/syahrul12345/Blocknalytics/node/modules/scraper/controllers"
)

func main() {
	go controllers.GetLatest()
	initRouter()
}

func initRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getAddress", controllers.GetTransactionsOfAccount).Methods("POST")
	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
