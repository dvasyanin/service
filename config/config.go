package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type App struct {
	Port string `yaml:"port"`
}

type Config struct {
	App *App `yaml:"app"`
}

func ReadFromFile(file string) (cfg *Config, err error) {
	cfgFile, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	cfg = &Config{}
	err = yaml.Unmarshal(cfgFile, cfg)
	return
}
