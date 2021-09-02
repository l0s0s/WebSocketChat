package client

import (
	"encoding/json"
	"io/ioutil"

	"go.uber.org/zap"
)

// Secret is google secret.
type Secret struct {
	Web map[string]interface{} `json:"web"`
}

// ReadJSON is read json fron file.
func ReadJSON(path string, logger *zap.Logger) map[string]interface{} {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatal("Failed to read json", zap.Error(err))
	}

	var config Secret

	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Fatal("Failed to unmarshal json", zap.Error(err))
	}
	return config.Web
}
