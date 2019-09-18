package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/syahrul12345/Blocknalytics/node/modules/analytics/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/getLatestBlock", controller.GetAccountInfo).Methods("POST")
	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
