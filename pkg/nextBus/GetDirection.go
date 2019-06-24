package nextBus

import (
	"encoding/json"
	"errors"
	"strconv"
)

func GetDirections(route int) (map[string]string, error) {
	if route < 0 {
		return nil, errors.New("invalid route number")
	}
	routeString := strconv.Itoa(route)
	js, err := GetEndpointData(buildDirectionsUrl(routeString))
	if err != nil {
		return nil, err
	}
	var textPairs []TextPairValue
	err = json.Unmarshal(js, &textPairs)
	if err != nil {
		return nil, err
	}
	directions := make(map[string]string)
	for _, v := range textPairs {
		directions[v.Text] = v.Value
	}
	return directions, nil
}
