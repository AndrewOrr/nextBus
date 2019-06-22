package nextBus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRoutes() []string {
	url := buildRoutesUrl()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("whoops")
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("hi")
		return nil
	}
	data, _ := ioutil.ReadAll(response.Body)
	var trips []NexTripRoute
	err = json.Unmarshal(data, &trips)
	if err != nil {
		fmt.Println("whoops")
		return nil
	}
	routes := make([]string, len(trips))
	for i, v := range trips {
		routes[i] = v.Description
	}
	return routes
}

type NexTripRoute struct {
	Description string `json:"Description"`
	ProviderId  string `json:"ProviderID"`
	Route       string `json:"Route"`
}
