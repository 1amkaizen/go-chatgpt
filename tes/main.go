package main

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 5,
		Prompt:    "how are you?",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
