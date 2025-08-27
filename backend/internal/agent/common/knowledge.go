package common

import (
	"context"
	"flowing/internal/model/chat"
	"flowing/internal/model/common"
	"flowing/internal/model/kb"
	"slices"

	"github.com/cloudwego/eino/schema"
)

// GetKnowledgeRefs 将知识库Retriever节点的输出，转换成知识库文档引用
func GetKnowledgeRefs(ctx context.Context, docs []*schema.Document) ([]chat.KnowledgeReference, error) {
	refs := make([]chat.KnowledgeReference, len(docs))
	// 转换成引用消息
	for i, doc := range docs {
		refs[i] = chat.KnowledgeReference{
			KnowledgeBaseId: doc.MetaData["knowledgeBaseId"].(int64),
			DocumentId:      doc.MetaData["docId"].(int64),
			SliceId:         doc.MetaData["sliceId"].(string),
			Content:         doc.Content,
		}
	}
	// 获取文档名称
	docNameMap, err := GetDocumentNames(ctx, refs)
	if err != nil {
		return nil, err
	}
	// 文档名称
	for i, ref := range refs {
		if name, ok := docNameMap[ref.DocumentId]; ok {
			refs[i].DocumentName = name
		}
	}
	return refs, nil
}

// GetKnowledgeRefMessage 将知识库Retriever节点的查询结果转换成发往用户的引用消息
func GetKnowledgeRefMessage(ctx context.Context, docs []*schema.Document, messageId,
	conversationId int64, agentId, agentRunId int64) (*chat.Message, error) {
	refs, err := GetKnowledgeRefs(ctx, docs)
	if err != nil {
		return nil, err
	}
	return &chat.Message{
		BaseModel:           common.BaseModel{Id: messageId},
		ConversationId:      conversationId,
		AgentId:             agentId,
		AgentRunId:          agentRunId,
		Type:                chat.MessageTypeAssistant,
		KnowledgeReferences: refs,
	}, nil
}

func GetDocumentNames(ctx context.Context, chunks []chat.KnowledgeReference) (map[int64]string, error) {
	// 获取文档名称
	docIds := make([]int64, 0)
	for _, chunk := range chunks {
		docIds = append(docIds, chunk.DocumentId)
	}
	// 去重id
	docIds = slices.Compact(docIds)
	docNames, err := kb.GetDocumentNames(ctx, docIds)
	if err != nil {
		return nil, err
	}
	// 文档名称map
	docNameMap := make(map[int64]string)
	for _, doc := range docNames {
		docNameMap[doc.Id] = doc.Name
	}
	return docNameMap, nil
}
