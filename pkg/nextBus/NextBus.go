package nextBus

import (
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
