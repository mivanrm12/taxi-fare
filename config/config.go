package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Fare FareConfig `yaml:"fare"`
}
type FareConfig struct {
	BaseDistanceThreshold int64 `yaml:"base_distance_threshold"`
	HighDistanceThreshold int64 `yaml:"high_distance_threshold"`
	BaseFare              int64 `yaml:"base_fare"`
	StandartRate          int64 `yaml:"standart_rate"`
	MediumDistanceRate    int64 `yaml:"medium_distance_rate"`
	HighDistanceRate      int64 `yaml:"high_distance_rate"`
}

func ReadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, err
}
