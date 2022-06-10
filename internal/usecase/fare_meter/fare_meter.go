package fare_meters

import (
	"github.com/mivanrm12/taxi-fare/config"
	"github.com/mivanrm12/taxi-fare/internal/model"
)

type FareMeter struct {
	fareCfg config.FareConfig
}

func New(cfg config.FareConfig) FareMeter {
	return FareMeter{
		fareCfg: cfg,
	}
}

func (fm FareMeter) CalculateTotalFare(input model.FareInput) model.FareOutput {
	if input.GetDistance() < float64(fm.fareCfg.BaseDistanceThreshold) {
		return model.FareOutput{
			Total: fm.fareCfg.BaseFare,
		}
	}
	total := fm.fareCfg.BaseFare + int64((input.Distance-float64(fm.fareCfg.BaseDistanceThreshold))/float64(fm.fareCfg.MediumDistanceRate))*fm.fareCfg.StandartRate

	if input.Distance > float64(fm.fareCfg.HighDistanceThreshold) {
		total += ((int64(input.Distance) - fm.fareCfg.HighDistanceThreshold) / fm.fareCfg.HighDistanceRate) * fm.fareCfg.StandartRate
	}
	return model.FareOutput{
		Total: total,
	}

}
