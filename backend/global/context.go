package global

type ContextKey string

var (
	ContextKeyUser      ContextKey = "user"
	ContextKeySSEWriter ContextKey = "SSEWriter"
)
