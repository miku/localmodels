package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
)

var (
	model         = flag.String("m", "zephyr", "model name for ollama")
	timeout       = flag.Duration("t", 120*time.Second, "timeout")
	systemMessage = flag.String("S", defaultSystemMessage, "system message to use")
	chatMessage   = flag.String("C", exampleMessage, "default chat message")

	defaultSystemMessage = `
Choose between two options A and B. You are presented with two options: A and B. Example input:

A

Gathered in syntax,
Go speaks with elegance and power,
A new world to explore.

B

a 5-7-5 rule
that's go's command line syntax
write your own stories /
your own rules for syntax.<|endoftext|>
`

	exampleMessage = `
A

To be true
Go runs on any platform
It’s portable and fast<|endoftext|>

B

Go is a modern language
Concise syntax, concurrency
Powerful, efficient
`
)

func main() {
	llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel(*model)))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), *timeout)
	defer cancelFunc()
	completion, err := llm.Call(ctx, []schema.ChatMessage{
		schema.SystemChatMessage{Content: *systemMessage},
		schema.HumanChatMessage{Content: *chatMessage},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		return nil
	}))
	if err != nil {
		log.Fatalf("%s failed: %v, skipping", model, err)
	}
	fmt.Println(completion.Content)
}
