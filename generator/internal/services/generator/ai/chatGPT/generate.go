package chatGPT

import (
	"context"

	"github.com/sashabaranov/go-openai"

	generatorModel "generator/internal/services/generator/model"

	"pkg/errors"
)

func (s *ChatGPTService) Generate(ctx context.Context, req generatorModel.GenerateReq) (res generatorModel.GenerateRes, err error) {

	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:         openai.ChatMessageRoleSystem,
				Content:      req.System,
				Refusal:      "",
				MultiContent: nil,
				Name:         "",
				FunctionCall: nil,
				ToolCalls:    nil,
				ToolCallID:   "",
			},
			{
				Role:         openai.ChatMessageRoleUser,
				Content:      req.Prompt,
				Refusal:      "",
				MultiContent: nil,
				Name:         "",
				FunctionCall: nil,
				ToolCalls:    nil,
				ToolCallID:   "",
			},
		},
		MaxTokens:           0,
		MaxCompletionTokens: 0,
		Temperature:         0,
		TopP:                0,
		N:                   0,
		Stream:              false,
		Stop:                nil,
		PresencePenalty:     0,
		ResponseFormat:      nil,
		Seed:                nil,
		FrequencyPenalty:    0,
		LogitBias:           nil,
		LogProbs:            false,
		TopLogProbs:         0,
		User:                "",
		Functions:           nil,
		FunctionCall:        nil,
		Tools:               nil,
		ToolChoice:          nil,
		StreamOptions:       nil,
		ParallelToolCalls:   nil,
		Store:               false,
		Metadata:            nil,
	})
	if err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	return generatorModel.GenerateRes{
		Text: resp.Choices[0].Message.Content,
	}, nil
}
