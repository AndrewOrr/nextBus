package nextBus

import (
	"reflect"
	"testing"
)

func TestGetDirections(t *testing.T) {
	type args struct {
		route int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"Basic test", args{6}, map[string]string{"NORTHBOUND": "4", "SOUTHBOUND": "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDirections(tt.args.route)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDirections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}
