package nextBus

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"time"
)

func GetNextTrip(route, direction, stop string) (string, error) {
	js, err := GetEndpointData(buildTimepointDepartureUrl(route, direction, stop))
	if err != nil {
		return "", err
	}
	var ntd []NexTripDeparture
	err = json.Unmarshal(js, &ntd)
	if err != nil {
		return "", err
	}
	r := regexp.MustCompile(`(?:\\/Date\\()(?P<date>[0-9]+)-(?P<tz>[0-9]+)(?:\\)/)`)
	//t := r.FindString(ntd[0].DepartureTime)
	matchArray := r.FindStringSubmatch(ntd[0].DepartureTime)
	if len(matchArray) < 3 {
		return "", errors.New("badly formated time")
	}
	ts, err := strconv.ParseInt(matchArray[1], 10, 64)
	if err != nil {
		return "", errors.New("badly formated time")
	}
	nextTime := time.Unix(int64(ts/1000), 0)
	currentTime := time.Now()
	eod := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 999999999, currentTime.Location())
	if nextTime.After(eod) {
		return "", nil
	}
	return ntd[0].DepartureText, nil
}
