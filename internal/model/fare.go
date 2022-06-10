package model

import "time"

type FareInput struct {
	TimeElapsed time.Time `json:"time_elapsed,omitempty"`
	Distance    float64   `json:"distance,omitempty"`
}

func (fi FareInput) GetTimeElapsed() time.Time {
	return fi.TimeElapsed
}
func (fi FareInput) GetDistance() float64 {
	return fi.Distance
}
func (fi *FareInput) SetTimeElapsed(timeElapsed time.Time) {
	fi.TimeElapsed = timeElapsed
}
func (fi *FareInput) SetDistance(distance float64) {
	fi.Distance = distance
}

type FareOutput struct {
	Total int64
}
