package nextBus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	url := buildRoutesUrl()
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
	var trips []NexTripRoute
	err = json.Unmarshal(data, &trips)
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

//NextTripRoute is the structure returned by metrotransit's GetRoutes
type NexTripRoute struct {
	Description string `json:"Description"`
	ProviderId  string `json:"ProviderID"`
	Route       string `json:"Route"`
}
