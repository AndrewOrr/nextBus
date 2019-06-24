package nextBus

import (
	"reflect"
	"testing"
	"time"
)

func TestGetNextTrip(t *testing.T) {
	type args struct {
		route     string
		direction string
		stop      string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"A known example should produce reasonable output", args{"6", "1", "1S1A"}, time.Now(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextTrip(tt.args.route, tt.args.direction, tt.args.stop)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNextTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
