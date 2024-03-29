package nextBus

import (
	"encoding/json"
	"errors"
	"strings"
)

func GetDirection(route, direction string) (string, error) {
	possibleDirections, err := GetDirections(route)
	if err != nil {
		return "", err
	}
	if val, ok := possibleDirections[strings.ToUpper(direction)]; ok {
		return val, nil
	}
	return "", errors.New("unknown direction")
}

func GetDirections(route string) (map[string]string, error) {
	js, err := GetEndpointData(buildDirectionsUrl(route))
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
