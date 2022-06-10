package main

import (
	"log"

	"github.com/mivanrm12/taxi-fare/config"
	"github.com/mivanrm12/taxi-fare/internal/handler"
	fare_meters "github.com/mivanrm12/taxi-fare/internal/usecase/fare_meter"
	"github.com/mivanrm12/taxi-fare/internal/usecase/validation"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting Taxi Fare App")
	config, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatal("failed to read Config", err)
	}

	validationService := validation.New()
	fareService := fare_meters.New(config.Fare)
	handler := handler.New(validationService, fareService)
	handler.HandleFunc()
}
