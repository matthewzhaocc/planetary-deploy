package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/matthewzhaocc/planetary-deploy/internal/types"
)

func ParseConfig(path string) (*types.ServerConfig, error) {
	var config types.ServerConfig

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
