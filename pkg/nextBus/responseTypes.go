package nextBus

import "encoding/json"

//NextTripRoute is the structure returned by metrotransit's GetRoutes
type NexTripRoute struct {
	Description string `json:"Description"`
	ProviderId  string `json:"ProviderID"`
	Route       string `json:"Route"`
}

type TextPairValue struct {
	Text  string
	Value string
}

type NexTripDeparture struct {
	Actual           bool
	BlockNumber      int
	DepartureText    string
	DepartureTime    string
	Description      string
	Gate             string
	Route            string
	RouteDirection   string
	Terminal         string
	VehicleHeading   int
	VehicleLatitude  json.Number
	VehicleLongitude json.Number
}
