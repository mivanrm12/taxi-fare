package handler

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mivanrm12/taxi-fare/internal/model"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)

	validationService := NewMockvalidationService(ctrl)
	fareService := NewMockfareService(ctrl)

	type args struct {
		validationService *MockvalidationService
		fareService       *MockfareService
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				validationService: validationService,
				fareService:       fareService,
			},
			want: &Handler{
				validationService: validationService,
				fareService:       fareService,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.validationService, tt.args.fareService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_HandleFunc(t *testing.T) {
	ctrl := gomock.NewController(t)

	ValidationService := NewMockvalidationService(ctrl)
	FareService := NewMockfareService(ctrl)
	type fields struct {
		validationService validationService
		fareService       fareService
	}
	tests := []struct {
		name   string
		fields fields
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				validationService: ValidationService,
				fareService:       FareService,
			},
			mock: func() {
				ValidationService.EXPECT().ValidateInput([]string{}).Return([]model.FareInput{
					{Distance: 100},
				}, nil)
				FareService.EXPECT().CalculateTotalFare(gomock.Any()).Return(model.FareOutput{})

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			h := Handler{
				validationService: tt.fields.validationService,
				fareService:       tt.fields.fareService,
			}
			h.HandleFunc()
		})
	}
}
