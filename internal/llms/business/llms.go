package business

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func (b *LLMsBusiness) Chat(question string) *string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 300,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
		Stream: true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	result := ""
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return &result
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return nil
		}

		fmt.Printf(response.Choices[0].Delta.Content)
		result += response.Choices[0].Delta.Content
	}
}
