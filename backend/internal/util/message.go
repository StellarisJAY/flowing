package util

import (
	"encoding/json"
	"flowing/internal/model/chat"
	"fmt"
	"io"
)

func SSESendMessage(message chat.Message, writer io.Writer) error {
	data, _ := json.Marshal(message)
	if _, err := fmt.Fprintf(writer, "data: %s\n\n", string(data)); err != nil {
		return err
	}
	return nil
}
