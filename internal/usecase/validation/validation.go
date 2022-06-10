package validation

import (
	"errors"
	"fmt"
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
		return []model.FareInput{}, errors.New("input length must be 2 or more line")
	}
	for i, input := range inputs {
		splittedInput := strings.Split(input, " ")
		if len(splittedInput) != 2 {
			logrus.Info(fmt.Sprintf("invalid input: %s, row: %d", input, i))
			return []model.FareInput{}, errors.New("input length must be formatted as \"timestamp distance\"")
		}
		timestamp, valid := validateTimestamp(splittedInput[0])
		if !valid {
			logrus.Info(fmt.Sprintf("invalid input: %s, row: %d", input, i))
			return []model.FareInput{}, errors.New("failed to parse timestamp input")
		}
		distance, valid := validateDistance(splittedInput[1])
		if !valid {
			logrus.Info(fmt.Sprintf("invalid input: %s, row: %d", input, i))
			return []model.FareInput{}, errors.New("failed to parse distance input")
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
		err := validatePrevious(fareInput, prevFare)
		if err != nil {
			return []model.FareInput{}, err
		}
		prevFare = fareInput
		fareInputs = append(fareInputs, fareInput)

	}
	return fareInputs, nil

}

func validateTimestamp(timestamp string) (time.Time, bool) {
	layout := "15:04:05.000"

	time, err := time.Parse(layout, timestamp)
	if err != nil {
		logrus.Info("invalid timestamp:", timestamp)
		return time, false
	}
	return time, true

}

func validateDistance(distance string) (float64, bool) {
	value, err := strconv.ParseFloat(strings.TrimSuffix(distance, "\n"), 64)
	if err != nil {
		logrus.Info("invalid distance: ", distance, err)
		return 0, false
	}
	return value, true

}

func validatePrevious(currentFare model.FareInput, prevFare model.FareInput) error {
	currentTime := currentFare.GetTimeElapsed()
	prevTime := prevFare.GetTimeElapsed()

	if currentTime.Before(prevTime) {
		return errors.New("past Time Already Sent")
	}
	if currentFare.Distance == 0 {
		return errors.New("distance shouldnt equal to 0")
	}
	if !prevTime.Add(time.Minute * 5).After(currentTime) {
		return errors.New("time differences greater than 5 minutes")
	}
	return nil

}
