package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigObj struct {
	Database struct {
		Path string `json:"path"`
	} `json:"database"`
}

func Load() ConfigObj {
	var config ConfigObj
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
