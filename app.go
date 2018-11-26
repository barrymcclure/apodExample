package main

import (
	"apodExample/apod_api"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var config appConfig

type serverAddress struct {
	IPAddress string `yaml:"Address"`
	Port      string `yaml:"Port"`
}

type appConfig struct {
	CacheServer serverAddress `yaml:"CacheServer"`
}

func main() {
	// read the configuration yaml
	configContent, err := ioutil.ReadFile("apod.yaml")
	if err != nil {
		//log the issue... continue without cache
		panic(err)
	}
	err = yaml.Unmarshal([]byte(configContent), &config)
	if err != nil {
		//
		panic(err)
	}

	cacheService := NewCacheServiceForApodClient(config.CacheServer.IPAddress, config.CacheServer.Port)
	apodClient := apod_api.NewApodServiceClient(cacheService)

	// site code is below...
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	t := "https://api.nasa.gov/planetary/apod?api_key=%s"
	apodServiceURL := fmt.Sprintf(t, os.Getenv("NASA_API"))
	fmt.Println(apodServiceURL)
	router.HandleFunc("/api/apod", apodClient.GetApod).Methods("GET")
	router.HandleFunc("/apodImg", apodClient.GetImgSource).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", router)
}

