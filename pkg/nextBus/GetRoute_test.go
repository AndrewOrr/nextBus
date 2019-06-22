package nextBus

import (
	"reflect"
	"testing"
)

func TestGetRoute(t *testing.T) {
	type args struct {
		routeInput string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Basic call should return expoected result", args{"nnepin - Xerxes - France - Southdale"}, 6, false},
		{"A string that doen't match should return an negative number", args{"qwert"}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoute(tt.args.routeInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRoutes(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoutes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoutes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
