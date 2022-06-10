package config

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				filename: "../config.yaml",
			},
			want: &Config{
				Fare: FareConfig{
					BaseDistanceThreshold: 1000,
					HighDistanceThreshold: 10000,
					BaseFare:              400,
					StandartRate:          40,
					MediumDistanceRate:    400,
					HighDistanceRate:      350,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
