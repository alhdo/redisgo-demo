package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var DataURL = "http://api.citybik.es/v2/networks/vlille"

type FetchHandler struct {
}

type Location struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Extra struct {
	Address         string `json:"address"`
	City            string `json:"city"`
	LastUpdate      string `json:"last_update"`
	Online          bool   `json:"online"`
	PaymentTerminal bool   `json:"payment-terminal"`
	Status          string `json:"status"`
	UID             string `json:"uid"`
}

type Station struct {
	EmptySlots int     `json:"empty_slots"`
	Extra      Extra   `json:"extra"`
	FreeBikes  int     `json:"free_bikes"`
	ID         string  `json:"id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Name       string  `json:"name"`
	Timestamp  string  `json:"timestamp"`
}

type Network struct {
	Company  []string  `json:"company"`
	Href     string    `json:"href"`
	ID       string    `json:"id"`
	Location Location  `json:"location"`
	Name     string    `json:"name"`
	Source   string    `json:"source"`
	Stations []Station `json:"stations"`
}

type NetworkResponse struct {
	Network Network `json:"network"`
}

func (fh *FetchHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	stations := FetchVelibData()

	b, err := json.Marshal(stations)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write(b); err != nil {
		fmt.Errorf("error writing response: %s", err)
	}
}

func FetchVelibData() []Station {
	response, err := http.Get(DataURL)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer response.Body.Close()
	// Decode the JSON response into a struct
	var networkResponse NetworkResponse
	if err := json.NewDecoder(response.Body).Decode(&networkResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return networkResponse.Network.Stations
}
