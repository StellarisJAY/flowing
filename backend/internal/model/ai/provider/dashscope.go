package provider

import "encoding/json"

type DashscopeConfig struct {
	ApiKey string `json:"apiKey"`
}

func DashscopeFromJSON(data []byte) (*DashscopeConfig, error) {
	var cfg DashscopeConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
