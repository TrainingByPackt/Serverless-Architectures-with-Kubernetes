package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
	"strconv"
	"time"
)

func FindBikes(w http.ResponseWriter, r *http.Request) {

	// Get query
	fmt.Println(r.URL.Query())
	query := r.URL.Query().Get("q")
	if query == "" {
		fmt.Println("No query is provided")
		w.Write([]byte(NO_BIKE_POINTS_FOUND))
		return
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}

	// Get bike points for the query
	bikePoints, err := httpClient.Get(fmt.Sprintf(TFL_API_URL + "BikePoint/Search?query=" + url2.QueryEscape(query)))
	if err != nil {
		fmt.Println("Error during request:", err)
		w.Write([]byte(NO_BIKE_POINTS_FOUND))
		return
	}
	defer r.Body.Close()

	// Parse data
	data := make([]BikePoint, 0)
	json.NewDecoder(bikePoints.Body).Decode(&data)
	if len(data) == 0 {
		w.Write([]byte(NO_BIKE_POINTS_FOUND))
		return
	}

	// Collect information
	bikePoint := data[0]
	url := fmt.Sprintf(GOOGLE_MAPS_URL_FORMAT, bikePoint.Lat, bikePoint.Lon)

	// Get available number of bikes
	availableBikeResponse, err := httpClient.Get(TFL_API_URL + "BikePoint/" + bikePoint.ID)
	if err != nil {
		fmt.Println("Error during request:", err)
		w.Write([]byte(fmt.Sprintf(RESPONSE_NO_BIKE_DATA, bikePoint.CommonName, url)))
		return
	}
	defer availableBikeResponse.Body.Close()

	// Parse data
	bikeData := Place{}
	json.NewDecoder(availableBikeResponse.Body).Decode(&bikeData)

	if len(bikeData.AdditionalProperties) == 0 {
		w.Write([]byte(fmt.Sprintf(RESPONSE_NO_BIKE_DATA, bikePoint.CommonName, url)))
		return

	}

	for _, property := range bikeData.AdditionalProperties {
		if property["key"] == "NbBikes" {
			bikeAmount, err := strconv.Atoi(property["value"])
			if err != nil {
				w.Write([]byte(fmt.Sprintf(RESPONSE_NO_BIKE_DATA, bikePoint.CommonName, url)))
				return
			}
			if bikeAmount == 0 {
				w.Write([]byte(fmt.Sprintf(RESPONSE_NO_AVAILABLE_BIKE, bikePoint.CommonName, url)))
				return
			} else {
				w.Write([]byte(fmt.Sprintf(DEFAULT_RESPONSE, bikePoint.CommonName, bikeAmount, url)))
				return
			}
		}
	}

}

const NO_BIKE_POINTS_FOUND = "Sorry, no bike points are found for this location!"
const DEFAULT_RESPONSE = "The nearest bike point is %s and there are %d available bikes 🚲 Hurry up! 🏃\nFor navigation 🧭 %s"
const RESPONSE_NO_AVAILABLE_BIKE = "The nearest bike point is %s but there are no available bikes ⚠️\nFor navigation 🧭 %s"
const RESPONSE_NO_BIKE_DATA = "The nearest bike point is %s but there are no information about bikes ⚠️\nFor navigation 🧭 %s"
const GOOGLE_MAPS_URL_FORMAT = "https://maps.google.com?q=%v,%v"
const TFL_API_URL = "https://api.tfl.gov.uk/"

type BikePoint struct {
	CommonName string
	Lat        float64
	Lon        float64
	ID         string
}

type Place struct {
	AdditionalProperties []map[string]string
}
