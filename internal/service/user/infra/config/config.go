package config

import (
	rc "woah/config"
	"woah/pkg/encoder/yaml"

	"github.com/tyr-tech-team/hawk/config"
)

// C -
var C = &Config{
	Info: config.Info{
		Name: "user",
	},
}

// Config -
type Config struct {
	Info    config.Info    `yaml:"info"`
	Mongo   config.Mongo   `yaml:"mongo"`
	Redis   config.Redis   `yaml:"redis"`
	Log     config.Log     `yaml:"log"`
	Service config.Service `yaml:"service"`
}

func NewServiceConfig(ic rc.IConfig) (*Config, error) {

	encoder := yaml.NewEncoder()
	err := encoder.Decode(ic.Get("services", C.Info.Name), &C)

	return C, err
}
