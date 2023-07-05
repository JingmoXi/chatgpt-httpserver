package chatgpt_api

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Call(msg string) (string, error) {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil
}
