package config

import (
	"woah/config"

	"woah/pkg/encoder/yaml"

	hconf "github.com/tyr-tech-team/hawk/config"
)

// C -
var C = Config{
	Info: hconf.Info{
		Name: "item",
	},
}

// Config -
type Config struct {
	Info    hconf.Info    `yaml:"info"`
	Mongo   hconf.Mongo   `yaml:"mongo"`
	Service hconf.Service `yaml:"service"`
	Mod     string        `yaml:"mod"`
	Log     hconf.Log     `yaml:"log"`
}

// New
func New(conf config.IConfig) (Config, error) {
	encoder := yaml.NewEncoder()

	err := encoder.Decode(conf.Get(config.SERVICES, C.Info.Name), &C)

	if C.Mod == "" {
		C.Mod = conf.Mod()
	}

	if C.Log.Level == "" {
		C.Log.Level = conf.Log()
	}

	if C.Mongo.Database == "" {
		C.Mongo = conf.MongoConf()
	}

	return C, err
}
