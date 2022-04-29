package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlFile struct {
	Config Config `yaml:"config"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func ReadConfig(path string) (*Config, error) {
	config := &YamlFile{}
	cfgFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(cfgFile, config)
	if err != nil {
		return nil, err
	}
	return &config.Config, nil
}
