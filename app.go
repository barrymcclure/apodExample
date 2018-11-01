package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var config appConfig

func main() {
	// read the configuration yaml
	configContent, err := ioutil.ReadFile("apod.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(configContent), &config)
	if err != nil {
		panic(err)
	}

	// site code is below...

	// router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static"))
	// t := "https://api.nasa.gov/planetary/apod?api_key=%s"
	// apodServiceURL := fmt.Sprintf(t, os.Getenv("NASA_API"))
	// fmt.Println(apodServiceURL)
	// router.HandleFunc("/api/apod", apod_api.GetApod).Methods("GET")
	// router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	// http.ListenAndServe(":8080", router)
}

type serverAddress struct {
	IPAddress string `yaml:"Address"`
	Port      string `yaml:"Port"`
}

type appConfig struct {
	CacheServer serverAddress `yaml:"CacheServer"`
}
