package mistral

import (
	"log"

	"github.com/gage-technologies/mistral-go"
)

func ChatStream(apiKey string, message string) {
	client := mistral.NewMistralClientDefault(apiKey)
	model := "codestral-latest"

	chatResChan, err := client.ChatStream(model, []mistral.ChatMessage{
		{Content: message, Role: mistral.RoleUser},
	}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}

	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", chatResChunk.Error)
		}
		log.Printf("Chat completion stream part: %+v\n", chatResChunk)
	}
}
