package util

import (
	"context"
	"flowing/internal/model/ai"
	"flowing/internal/model/ai/provider"
	"fmt"

	"github.com/cloudwego/eino-ext/components/embedding/dashscope"
	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino/components/embedding"
)

func GetEmbeddingModel(ctx context.Context, pm ai.ProviderModelDetail) (embedding.Embedder, error) {
	cfg, err := ai.GetProviderConfig(pm.ProviderType, pm.ProviderConfig)
	if err != nil {
		return nil, err
	}
	var embedder embedding.Embedder
	switch pm.ProviderType {
	case ai.ProviderTypeOpenAI:
		openAICfg := cfg.(*provider.OpenAIProviderConfig)
		embedder, err = openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
			Model:   pm.ModelName,
			BaseURL: openAICfg.BaseUrl,
			APIKey:  openAICfg.ApiKey,
		})
	case ai.ProviderTypeDashscope:
		dashscopeCfg := cfg.(*provider.DashscopeConfig)
		embedder, err = dashscope.NewEmbedder(ctx, &dashscope.EmbeddingConfig{
			APIKey: dashscopeCfg.ApiKey,
			Model:  pm.ModelName,
		})
	default:
		return nil, fmt.Errorf("unsupported provider type: %v", pm.ProviderType)
	}
	if err != nil {
		return nil, err
	}
	return embedder, nil
}

func Float64EmbeddingTo32(embedding []float64) []float32 {
	float32Embedding := make([]float32, len(embedding))
	for i, v := range embedding {
		float32Embedding[i] = float32(v)
	}
	return float32Embedding
}
