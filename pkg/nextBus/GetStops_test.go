package nextBus

import (
	"reflect"
	"testing"
)

func TestGetStopIdentifier(t *testing.T) {
	type args struct {
		route     string
		direction string
		stop      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"random input should return an error", args{"aaa", "aaa", "aaa"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStopIdentifier(tt.args.route, tt.args.direction, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStopIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStopIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStops(t *testing.T) {
	type args struct {
		route     string
		direction string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		// TODO: this isn't a consistent test
		{"given a known route, should return expected results", args{"6", "1"}, nil, true},
		{"random input should return an error", args{"aaa", "aaa"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStops(tt.args.route, tt.args.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStops() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
