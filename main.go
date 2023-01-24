package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	var looping = true

	for looping {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your prompt: ")
		input, _ := reader.ReadString('\n')

		c := gogpt.NewClient("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")
		ctx := context.Background()

		req := gogpt.CompletionRequest{
			Model: gogpt.GPT3TextDavinci002,
			//			MaxTokens: 9999,
			Prompt: input,
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return
		}
		fmt.Println(resp.Choices[0].Text)
	}
}
