package util

import (
	"context"
	"errors"
	"flowing/internal/model/ai"
	"flowing/internal/model/ai/provider"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/components/model"
)

func GetChatModel(ctx context.Context, pm ai.ProviderModelDetail, outputJson bool) (model.BaseChatModel, error) {
	var cm model.BaseChatModel
	var err error
	format := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeText,
	}
	if outputJson {
		format.Type = openai.ChatCompletionResponseFormatTypeJSONObject
	}
	cnf, err := ai.GetProviderConfig(pm.ProviderType, pm.ProviderConfig)
	if err != nil {
		return nil, errors.New("invalid model config")
	}
	switch pm.ProviderType {
	case ai.ProviderTypeOpenAI:
		config := cnf.(*provider.OpenAIProviderConfig)

		cm, err = openai.NewChatModel(ctx, &openai.ChatModelConfig{
			APIKey:         config.ApiKey,
			BaseURL:        config.BaseUrl,
			Model:          pm.ModelName,
			ResponseFormat: &format,
		})
	case ai.ProviderTypeDashscope:
		config := cnf.(*provider.DashscopeConfig)
		enableThinking := true
		cm, err = qwen.NewChatModel(ctx, &qwen.ChatModelConfig{
			BaseURL:        "https://dashscope.aliyuncs.com/compatible-mode/v1",
			APIKey:         config.ApiKey,
			Model:          pm.ModelName,
			EnableThinking: &enableThinking,
			ResponseFormat: &format,
		})
	default:
		return nil, errors.New("unsupported provider type")
	}
	if err != nil {
		return nil, err
	}
	return cm, nil
}
