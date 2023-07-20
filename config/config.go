package config

import (
	"log"
	"main/model"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadData() *model.ConfigData {
	configData := &model.ConfigData{}

	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}

	err = yaml.Unmarshal(yamlFile, configData)
	if err != nil {
		log.Fatal("Failed to map config data:", err)
	}

	return configData
}
