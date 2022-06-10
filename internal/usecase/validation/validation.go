package validation

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/mivanrm12/taxi-fare/internal/model"

	"github.com/sirupsen/logrus"
)

type Validator struct{}

func New() Validator {
	return Validator{}
}
func (v Validator) ValidateInput(inputs []string) ([]model.FareInput, error) {
	var (
		prevFare model.FareInput
	)
	fareInputs := []model.FareInput{}
	if len(inputs) < 2 {
		logrus.Warn("Input length <=2")
		return fareInputs, errors.New("input length must be 2 or more line")
	}
	for i, input := range inputs {
		splittedInput := strings.Split(input, " ")
		if len(splittedInput) != 2 {
			logrus.Info("invalid input: %s, row: %d", input, i)
			return fareInputs, errors.New("input length must be formatted as \"timestamp distance\"")
		}
		timestamp, valid := validateTimestamp(splittedInput[0])
		if !valid {
			logrus.Info("invalid input: %s, row: %d", input, i)
			return fareInputs, errors.New("failed to parse timestamp input")
		}
		distance, valid := validateDistance(splittedInput[1])
		if !valid {
			logrus.Info("invalid input: %s, row: %d", input, i)
			return fareInputs, errors.New("failed to parse distance input")
		}
		fareInput := model.FareInput{
			TimeElapsed: timestamp,
			Distance:    distance,
		}
		if i == 0 {
			prevFare = fareInput
			fareInputs = append(fareInputs, fareInput)
			continue
		}
		if validatePrevious(fareInput, prevFare) {
			prevFare = fareInput
			fareInputs = append(fareInputs, fareInput)
		}

	}
	return fareInputs, nil

}

func validateTimestamp(timestamp string) (time.Time, bool) {
	layout := "15:04:05.000"

	time, err := time.Parse(layout, timestamp)
	if err != nil {
		logrus.Info("invalid timestamp: %s", timestamp)
		return time, false
	}
	return time, true

}

func validateDistance(distance string) (float64, bool) {
	value, err := strconv.ParseFloat(strings.TrimSuffix(distance, "\n"), 64)
	if err != nil {
		logrus.Info("invalid distance: %s", distance, err)
		return 0, false
	}
	return value, true

}

func validatePrevious(currentFare model.FareInput, prevFare model.FareInput) bool {
	currentTime := currentFare.GetTimeElapsed()
	prevTime := prevFare.GetTimeElapsed()

	if currentTime.Before(prevTime) {
		return false
	}
	if currentFare.Distance == 0 {
		return false
	}
	return prevTime.Add(time.Minute * 5).After(currentTime)
}
