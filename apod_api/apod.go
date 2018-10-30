package apod_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// formatted string... to pass to printf or what have you
var apodServiceUrl = "https://api.nasa.gov/planetary/apod?api_key=%s"

type ApodResponse struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func fetch() ApodResponse {
	c := http.Client{Timeout: time.Second * 60}
	response, err := c.Get(fmt.Sprintf(apodServiceUrl, os.Getenv("NASA_API")))
	if err != nil {
		fmt.Printf("There was an error, %v", err)
	}
	//fmt.Printf("Content-Type is %s.", response.Header.Get("Content-Type"))
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("There was an error, %v", err)
	}
	var apod ApodResponse
	json.Unmarshal(b, &apod)
	response.Body.Close()
	return apod
}

func GetApod(w http.ResponseWriter, h *http.Request) {
	apod := fetch()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apod)
}
