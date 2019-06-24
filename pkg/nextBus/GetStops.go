package nextBus

import (
	"encoding/json"
	"errors"
)

func GetStopIdentifier(route, direction, stop string) (string, error) {
	mapping, err := GetStops(route, direction)
	if err != nil {
		return "", err
	}
	for k, v := range mapping {
		if k == stop {
			return v, nil
		}
	}
	return "", errors.New("unknown stop value")
}

func GetStops(route, direction string) (map[string]string, error) {
	js, err := GetEndpointData(buildStopUrl(route, direction))
	if err != nil {
		return nil, err
	}
	var stops []TextPairValue
	err = json.Unmarshal(js, &stops)
	if err != nil {
		return nil, err
	}
	stopToIdentifier := make(map[string]string)
	for _, v := range stops {
		stopToIdentifier[v.Text] = v.Value
	}
	return stopToIdentifier, nil
}
