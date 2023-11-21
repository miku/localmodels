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
We are in an exam situation.
`

	exampleMessage = `

You are driving towards a traffic light which is switched to red, what do you do?

A: stop
B: continue driving
`
)

func main() {
	flag.Parse()
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
