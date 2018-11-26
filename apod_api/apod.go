package apod_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ApodCacheService interface {
	Get(key string, callback func() string ) string
}

type ApodServiceClient struct {
	cacheService ApodCacheService
}

// formatted string... to pass to printf or what have you
var apodServiceUrl = "https://api.nasa.gov/planetary/apod?api_key=%s"

func fetch() string {
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
	response.Body.Close()
	return string(b)
}

func NewApodServiceClient(cache ApodCacheService) *ApodServiceClient {
	asc := ApodServiceClient{
		cacheService: cache,
	}
	return &asc
}

func  (client *ApodServiceClient ) GetApod(w http.ResponseWriter, h *http.Request) {
	//fetch data from cache
	//c := cacheService.NewCacher(config.CacheServer.IPAddress, config.CacheServer.Port)
	apod := client.cacheService.Get("apodData", fetch)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(apod))
	// this writes to the response writer the json data in variable apod.
	//json.NewEncoder(w).Encode(apod)
}


func (client *ApodServiceClient) GetImgSource(w http.ResponseWriter, h *http.Request) {

}

