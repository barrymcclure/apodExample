package main

import (
	"apodExample/apod_api"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	t := "https://api.nasa.gov/planetary/apod?api_key=%s"
	apodServiceURL := fmt.Sprintf(t, os.Getenv("NASA_API"))
	fmt.Println(apodServiceURL)
	router.HandleFunc("/api/apod", apod_api.GetApod).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", router)
}
