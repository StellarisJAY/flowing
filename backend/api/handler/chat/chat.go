package chat

import (
	"context"
	"flowing/global"
	"flowing/internal/messagehub"
	"flowing/internal/model/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	var req chat.SendMessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	writer := c.Writer
	writer.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")
	_, ok := writer.(http.Flusher)
	if !ok {
		panic(global.NewError(500, "浏览器不支持SSE", nil))
	}
	ctx := context.WithValue(c.Request.Context(), global.ContextKeySSEWriter, writer)
	err := messagehub.HandleSendMessage(ctx, req)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
