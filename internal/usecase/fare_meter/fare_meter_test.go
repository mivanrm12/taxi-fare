package fare_meters

import (
	"reflect"
	"testing"
	"time"

	"github.com/mivanrm12/taxi-fare/config"
	"github.com/mivanrm12/taxi-fare/internal/model"
)

func TestFareMeter_CalculateTotalFare(t *testing.T) {
	type fields struct {
		fareCfg config.FareConfig
	}
	type args struct {
		input model.FareInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.FareOutput
	}{
		{
			name: "base fare",
			fields: fields{
				fareCfg: config.FareConfig{
					BaseDistanceThreshold: 1000,
					BaseFare:              400,
				},
			},
			args: args{
				input: model.FareInput{
					TimeElapsed: time.Now(),
					Distance:    100,
				},
			},
			want: model.FareOutput{
				Total: 400,
			},
		},
		{
			name: "medium distance",
			fields: fields{
				fareCfg: config.FareConfig{
					BaseDistanceThreshold: 1000,
					HighDistanceThreshold: 10000,
					BaseFare:              400,
					StandartRate:          40,
					MediumDistanceRate:    400,
					HighDistanceRate:      350,
				},
			},
			args: args{
				input: model.FareInput{
					TimeElapsed: time.Now(),
					Distance:    1500,
				},
			},
			want: model.FareOutput{
				Total: 440,
			},
		},
		{
			name: "high distance",
			fields: fields{
				fareCfg: config.FareConfig{
					BaseDistanceThreshold: 1000,
					HighDistanceThreshold: 10000,
					BaseFare:              400,
					StandartRate:          40,
					MediumDistanceRate:    400,
					HighDistanceRate:      350,
				},
			},
			args: args{
				input: model.FareInput{
					TimeElapsed: time.Now(),
					Distance:    10500,
				},
			},
			want: model.FareOutput{
				Total: 1360,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := FareMeter{
				fareCfg: tt.fields.fareCfg,
			}
			if got := fm.CalculateTotalFare(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FareMeter.CalculateTotalFare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cfg config.FareConfig
	}
	tests := []struct {
		name string
		args args
		want FareMeter
	}{
		{
			name: "success",
			args: args{
				cfg: config.FareConfig{},
			},
			want: FareMeter{
				fareCfg: config.FareConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
