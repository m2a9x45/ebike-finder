package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Bikes struct {
	Data struct {
		StationMapPage struct {
			Title    string      `json:"title"`
			Heading  interface{} `json:"heading"`
			Typename string      `json:"__typename"`
		} `json:"stationMapPage"`
		System struct {
			InSeason bool `json:"inSeason"`
			Coords   []struct {
				Lat      float64 `json:"lat"`
				Lng      float64 `json:"lng"`
				Typename string  `json:"__typename"`
			} `json:"coords"`
			Typename string `json:"__typename"`
		} `json:"system"`
		DockGroups []struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			Title            string `json:"title"`
			State            string `json:"state"`
			SubTitle         string `json:"subTitle"`
			AvailabilityInfo struct {
				AvailableVehicles          int `json:"availableVehicles"`
				AvailableDocks             int `json:"availableDocks"`
				AvailableVehicleCategories []struct {
					Category string `json:"category"`
					Count    int    `json:"count"`
					Typename string `json:"__typename"`
				} `json:"availableVehicleCategories"`
				Typename string `json:"__typename"`
			} `json:"availabilityInfo"`
			Coord struct {
				Lat      float64 `json:"lat"`
				Lng      float64 `json:"lng"`
				Typename string  `json:"__typename"`
			} `json:"coord"`
			Typename string `json:"__typename"`
		} `json:"dockGroups"`
	} `json:"data"`
}

func main() {
	http.HandleFunc("/", handler)

	r := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin", "Access-Control-Request-Headers"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})
	origin := handlers.AllowedOrigins([]string{"http://127.0.0.1:8080", "http://bike.lewisdukelow.com"})
	creds := handlers.AllowCredentials()

	r.HandleFunc("/", handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":7000", handlers.CORS(header, methods, origin, creds)(r)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := "https://core.urbansharing.com/public/api/v1/graphql?operationName=stationMapQuery&variables=%7B%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%225e1487fc64b3f3418e57abeee08c7c7aae79dd42d62c278500a698731dc6642a%22%7D%7D"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("systemid", "edinburgh-city-bikes")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bikeData := Bikes{}

	jsn, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error wilst reading r body", err)
	}

	err = json.Unmarshal(jsn, &bikeData)
	if err != nil {
		log.Fatal("Error wilst unmarshaling json", err)
	}

	// log.Println(bikeData)

	json.NewEncoder(w).Encode(bikeData)

}
