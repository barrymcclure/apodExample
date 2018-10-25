package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	t := "https://api.nasa.gov/planetary/apod?api_key=%s"
	apodServiceUrl := fmt.Sprintf(t, os.Getenv("NASA_API"))
	fmt.Println(apodServiceUrl)

	//mux.HandleFunc("/", greet)
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
