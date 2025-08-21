package docprocess

import (
	"context"
	"errors"
	"flowing/internal/model/kb"
	"io"

	"github.com/cloudwego/eino-ext/components/document/parser/pdf"
	"github.com/cloudwego/eino/components/document/parser"
	"github.com/cloudwego/eino/schema"
)

func ParseDocument(ctx context.Context, doc *kb.Document, content io.Reader) ([]*schema.Document, error) {
	var docs []*schema.Document
	var err error
	// 读取PDF文件
	if doc.Type == "pdf" {
		if doc.MIMEType != "application/pdf" {
			return nil, errors.New("not pdf")
		}
		pdfParser, err := pdf.NewPDFParser(ctx, &pdf.Config{ToPages: false})
		if err != nil {
			return nil, err
		}
		docs, err = pdfParser.Parse(ctx, content)
		if err != nil {
			return nil, err
		}
	} else if doc.Type == "md" || doc.Type == "txt" { // 读取文本文件
		docs, err = parser.TextParser{}.Parse(ctx, content)
		if err != nil {
			return nil, err
		}
	}
	return docs, nil
}
