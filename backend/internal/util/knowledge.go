package util

import "fmt"

const (
	CollectionNamePrefix = "flowingkb"
)

func KnowledgeBaseCollectionName(id int64) string {
	return fmt.Sprintf("%s%d", CollectionNamePrefix, id)
}
