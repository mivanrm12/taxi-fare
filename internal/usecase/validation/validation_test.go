package validation

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/mivanrm12/taxi-fare/internal/model"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Validator
	}{
		{
			name: "success",
			want: Validator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_ValidateInput(t *testing.T) {
	layout := "15:04:05.000"

	time1, _ := time.Parse(layout, "00:00:00.000")
	time2, _ := time.Parse(layout, "00:01:00.123")

	type args struct {
		inputs []string
	}
	tests := []struct {
		name    string
		v       Validator
		args    args
		want    []model.FareInput
		wantErr bool
	}{
		{
			name: "input less than 2",
			v:    Validator{},
			args: args{
				inputs: []string{},
			},
			want:    []model.FareInput{},
			wantErr: true,
		},
		{
			name: "input less than 2",
			v:    Validator{},
			args: args{
				inputs: []string{},
			},
			want:    []model.FareInput{},
			wantErr: true,
		},
		{
			name: "input splitted into more than 2",
			v:    Validator{},
			args: args{
				inputs: []string{
					"00:00:00 1.5 1",
					"00:00:00 1.5 1",
				},
			},
			want:    []model.FareInput{},
			wantErr: true,
		},
		{
			name: "failed parse timestamp",
			v:    Validator{},
			args: args{
				inputs: []string{
					"00:0000.000 0.0",
					"00:01:00.123 480.9",
				},
			},
			want:    []model.FareInput{},
			wantErr: true,
		}, {
			name: "failed parse distance",
			v:    Validator{},
			args: args{
				inputs: []string{
					"00:00:00.000 0.0.0",
					"00:01:00.123 480.9",
				},
			},
			want:    []model.FareInput{},
			wantErr: true,
		},
		{
			name: "prev time already sent",
			v:    Validator{},
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
					"00:00:50.123 480.9",
				},
			},
			want:    []model.FareInput{},
			wantErr: true,
		},
		{
			name: "success",
			v:    Validator{},
			args: args{
				inputs: []string{
					"00:00:00.000 0.0",
					"00:01:00.123 480.9",
				},
			},
			want: []model.FareInput{
				{
					TimeElapsed: time1,
					Distance:    0,
				},
				{
					TimeElapsed: time2,
					Distance:    480.9,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{}
			got, err := v.ValidateInput(tt.args.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validator.ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateTimestamp(t *testing.T) {
	type args struct {
		timestamp string
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateTimestamp(tt.args.timestamp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateTimestamp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("validateTimestamp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_validateDistance(t *testing.T) {

	type args struct {
		distance string
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateDistance(tt.args.distance)
			if got != tt.want {
				t.Errorf("validateDistance() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("validateDistance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_validatePrevious(t *testing.T) {
	layout := "15:04:05.000"

	time1, _ := time.Parse(layout, "00:00:00.000")
	time2, _ := time.Parse(layout, "00:01:00.123")
	time3, _ := time.Parse(layout, "00:09:00.123")
	type args struct {
		currentFare model.FareInput
		prevFare    model.FareInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "past time",
			args: args{
				currentFare: model.FareInput{
					TimeElapsed: time1,
				},
				prevFare: model.FareInput{
					TimeElapsed: time2,
				},
			},
			wantErr: true,
		},
		{
			name: "distance 0 ",
			args: args{
				currentFare: model.FareInput{
					TimeElapsed: time3,
					Distance:    0,
				},
				prevFare: model.FareInput{
					TimeElapsed: time1,
				},
			},
			wantErr: true,
		},
		{
			name: "5 minutes gap",
			args: args{
				currentFare: model.FareInput{
					TimeElapsed: time3,
					Distance:    500,
				},
				prevFare: model.FareInput{
					TimeElapsed: time1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePrevious(tt.args.currentFare, tt.args.prevFare); (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("validatePrevious() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
