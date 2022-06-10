package handler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mivanrm12/taxi-fare/internal/model"

	"github.com/sirupsen/logrus"
)

type validationService interface {
	ValidateInput([]string) ([]model.FareInput, error)
}

type fareService interface {
	CalculateTotalFare(input model.FareInput) (model.FareOutput, error)
}
type Handler struct {
	validationService validationService
	fareService       fareService
}

func New(validationService validationService, fareService fareService) *Handler {
	return &Handler{
		validationService: validationService,
		fareService:       fareService,
	}
}

func (h Handler) HandleFunc() {
	var currentFare model.FareOutput
	reader := bufio.NewReader(os.Stdin)
	inputs := []string{}
	fmt.Println("Please Input with \"timestamp distance\" format, program will stop receiving input if provided with blank line")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			logrus.Error(err)
			break
		}
		if text == "\n" {
			break
		}
		inputs = append(inputs, text)
	}

	//validate input
	fareInput, err := h.validationService.ValidateInput(inputs)
	if err != nil {
		logrus.Error(err)
	}

	//calculate fare
	for _, input := range fareInput {

		currentFare, err = h.fareService.CalculateTotalFare(input)
		if err != nil {
			logrus.Error(err)
		}

	}
	fmt.Println(currentFare.Total)

}
