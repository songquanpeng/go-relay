package common

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
}

var CONFIG = Config{}

func initConfigFile() {
	if _, err := os.Stat(*ConfigFile); err == nil {
		println("Config file already exists.")
		os.Exit(1)
	}
	defaultConfig := Config{
		Port:  6972,
		Token: GenerateToken(),
	}
	defaultConfigBytes, err := yaml.Marshal(defaultConfig)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = os.WriteFile(*ConfigFile, defaultConfigBytes, 0644)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Initialized config file at: " + *ConfigFile)
}

func loadConfigFile() {
	// Check if config file exists.
	if _, err := os.Stat(*ConfigFile); err != nil {
		initConfigFile()
	}
	configBytes, err := os.ReadFile(*ConfigFile)
	if err != nil {
		println("Failed to read config file: " + err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(configBytes, &CONFIG)
	if err != nil {
		println("Failed to parse config file: " + err.Error())
		os.Exit(1)
	}
}
