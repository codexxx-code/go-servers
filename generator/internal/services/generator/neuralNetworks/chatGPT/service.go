package chatGPT

import "github.com/sashabaranov/go-openai"

type ChatGPTService struct {
	client *openai.Client
}

func NewChatGPTService(client *openai.Client) *ChatGPTService {
	return &ChatGPTService{
		client: client,
	}
}
