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
	AgeLimit struct {
		Min int `json:min`
		Max int `json:max`
	} `json:ageLimit`
	AmountLimit struct {
		Min        int `json:min`
		Max        int `json:max`
		Multiplier int `json:multiplier`
	} `json:amountLimit`
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
