// mistral/mistral.go
package mistral

import (
	"log"
	"strings"
	"sync"

	"github.com/gage-technologies/mistral-go"
)

func ChatStream(apiKey, message string, response *strings.Builder, mu *sync.Mutex) {
	client := mistral.NewMistralClientDefault(apiKey)
	model := "codestral-latest"

	chatResChan, err := client.ChatStream(
		model,
		[]mistral.ChatMessage{
			{Content: message, Role: mistral.RoleUser},
		},
		nil,
	)
	if err != nil {
		log.Printf("Error getting chat completion stream: %v", err)
		return
	}

	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Printf("Error while streaming response: %v", chatResChunk.Error)
			return
		}

		for _, choice := range chatResChunk.Choices {
			content := choice.Delta.Content
			if content != "" {
				mu.Lock()
				response.WriteString(content)
				mu.Unlock()
			}
		}
	}
}
