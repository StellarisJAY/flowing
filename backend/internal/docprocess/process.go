package docprocess

import (
	"context"
	"flowing/internal/model/kb"
	"flowing/internal/repository"
	"io"
)

type Document struct {
	Content io.Reader
	Model   *kb.Document
}

type Option struct {
	EmbedOption
	SplitOption
}

var SupportedTypes = []string{"pdf", "txt", "md"}

func generateSliceId(ctx context.Context, originalID string, splitIndex int) string {
	return repository.Snowflake().Generate().String()
}
