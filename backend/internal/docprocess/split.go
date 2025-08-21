package docprocess

import (
	"context"
	"flowing/internal/model/kb"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/recursive"
	"github.com/cloudwego/eino/schema"
)

type SplitOption struct {
	ChunkSize   int
	OverlapSize int
	Separators  []string
}

func splitMarkdown(ctx context.Context, docs []*schema.Document) ([]*schema.Document, error) {
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "title",
			"##":  "section",
			"###": "subsection",
		},
		TrimHeaders: false,
		IDGenerator: generateSliceId,
	})
	if err != nil {
		return nil, err
	}
	return splitter.Transform(ctx, docs)
}

func splitTextRecursive(ctx context.Context, docs []*schema.Document, opt *SplitOption) ([]*schema.Document, error) {
	splitter, err := recursive.NewSplitter(ctx, &recursive.Config{
		ChunkSize:   opt.ChunkSize,
		OverlapSize: opt.OverlapSize,
		Separators:  opt.Separators,
		IDGenerator: generateSliceId,
	})
	if err != nil {
		return nil, err
	}
	return splitter.Transform(ctx, docs)
}

func SplitDocument(ctx context.Context, docs []*schema.Document, originalDoc *kb.Document, opt *SplitOption) ([]*schema.Document, error) {
	var chunks []*schema.Document
	var err error
	chunks, err = splitTextRecursive(ctx, docs, opt)
	//switch originalDoc.Type {
	//case "md":
	//	chunks, err = splitMarkdown(ctx, docs)
	//default:
	//	chunks, err = splitTextRecursive(ctx, docs, opt)
	//}
	return chunks, err
}
