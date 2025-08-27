package docprocess

import (
	"context"
	"flowing/internal/model/ai"
	"flowing/internal/util"

	"github.com/cloudwego/eino/schema"
)

type EmbedOption struct {
	ProviderType ai.ProviderType
	ModelName    string
	Config       string
}

const embedBatchSize = 10

func EmbedQuery(ctx context.Context, query string, opt *EmbedOption) ([]float64, error) {
	embeddings, err := Embed(ctx, []*schema.Document{{Content: query}}, opt)
	if err != nil {
		return nil, err
	}
	return embeddings[0], nil
}

func Embed(ctx context.Context, docs []*schema.Document, opt *EmbedOption) ([][]float64, error) {
	texts := make([]string, len(docs))
	for i, doc := range docs {
		texts[i] = doc.Content
	}
	embedder, err := util.GetEmbeddingModel(ctx, ai.ProviderModelDetail{
		ProviderType:   opt.ProviderType,
		ModelName:      opt.ModelName,
		ProviderConfig: opt.Config,
	})
	if err != nil {
		return nil, err
	}

	embeddings := make([][]float64, 0, len(texts))
	for i := 0; i < len(texts); i += embedBatchSize {
		end := min(i+embedBatchSize, len(texts))
		emb, err := embedder.EmbedStrings(ctx, texts[i:end])
		if err != nil {
			return nil, err
		}
		embeddings = append(embeddings, emb...)
	}
	return embeddings, nil
}
