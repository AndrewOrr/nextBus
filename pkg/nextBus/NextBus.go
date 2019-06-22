package nextBus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	baseUrl = "http://svc.metrotransit.org/NexTrip/"
)

func buildRoutesUrl() string {
	return baseUrl + "Routes"
}

func buildDirectionsUrl(route string) string {
	return baseUrl + "Directions/" + route
}

func buildStopUrl(route string, direction string) string {
	return baseUrl + "Stops/" + route + "/" + direction
}

func buildAllDeparturesUrl(stopId int) string {
	return baseUrl + strconv.Itoa(stopId)
}

func buildTimepointDepartureUrl(route string, direction string, stop string) string {
	return baseUrl + route + "/" + direction + "/" + stop
}


//GetEndpointData handles calls to the metrotransit api
func GetEndpointData(url string) (json.RawMessage, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}
