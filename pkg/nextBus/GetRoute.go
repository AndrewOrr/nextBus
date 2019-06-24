package nextBus

import (
	"encoding/json"
	"strconv"
	"strings"
)

//GetRoute take in a string and compares it against the list of routes. If a match is found it returns the route #
//if it fails to find it returns -1
func GetRoute(routeInput string) (int, error) {
	routes, err := GetRoutes()
	if err != nil {
		return 0, err
	}
	//ToDo: this is pretty costly
	for i, v := range routes {
		if strings.Contains(i, routeInput) {
			return v, nil
		}
	}
	return -1, nil
}

//GetRoutes calles the metrotransit api and returns a map of descripts -> route number
func GetRoutes() (map[string]int, error) {
	js, err := GetEndpointData(buildRoutesUrl())
	if err != nil {
		return nil, err
	}
	var trips []NexTripRoute
	err = json.Unmarshal(js, &trips)
	if err != nil {
		return nil, err
	}
	routes := make(map[string]int)
	for _, v := range trips {
		routes[v.Description], err = strconv.Atoi(v.Route)
		if err != nil {
			return nil, err
		}
	}
	return routes, nil
}
