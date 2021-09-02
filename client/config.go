package client

import (
	"encoding/json"
	"io/ioutil"

	"go.uber.org/zap"
)

type ClientSecret struct {
	Web map[string]interface{} `json:"web"`
}

func ReadJson(path string, logger *zap.Logger) map[string]interface{} {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatal("Failed to read \"reddit_jokes.json\"", zap.Error(err))
	}

	var config ClientSecret

	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Fatal("Failed to unmarshal json", zap.Error(err))
	}
	return config.Web
}