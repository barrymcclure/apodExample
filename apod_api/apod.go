package apod_api

import (
	"net/http"
	"time"
	"fmt"
	"os"
	"io/ioutil"
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

func fetch () {
	c := http.Client{Timeout:time.Second*60}
	response, err := c.Get(fmt.Sprintf(apodServiceUrl, os.Getenv("NASA_API")))
	if err != nil {
		fmt.Printf("There was an error, %v", err)
	}
	//fmt.Printf("Content-Type is %s.", response.Header.Get("Content-Type"))
	ioutil.ReadAll(response.Body)
	response.Body.Close()

}
