package config

import (
	 "encoding/json"
	 "os"
)

type Configuration struct {
    LogToCloud    bool
    LogToLocal	  bool
}


func LoadConfig(filePath string) (Configuration, error) {
	file, _ := os.Open(filePath)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	return configuration,err
}