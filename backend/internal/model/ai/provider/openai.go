package provider

import "encoding/json"

type OpenAIProviderConfig struct {
	BaseUrl string `json:"baseUrl"`
	ApiKey  string `json:"apiKey"`
}

func OpenAIFromJSON(data []byte) (*OpenAIProviderConfig, error) {
	var cfg OpenAIProviderConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
