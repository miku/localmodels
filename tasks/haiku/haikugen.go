package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
)

var (
	availableModels = []string{
		"llama2",
		"zephyr",
		"mistral",
		"mistrallite",
		"falcon",
		"orca-mini",
	}
	numSamples  = flag.Int("n", 1, "number of samples to generate")
	outputStyle = flag.String("t", "", "output mode, leave empty for stdout, or json for structure")
)

func main() {
	flag.Parse()
	for _, model := range availableModels {
		llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel(model)))
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		for i := 0; i < *numSamples; i++ {
			completion, err := llm.Call(ctx, []schema.ChatMessage{
				schema.SystemChatMessage{Content: "Task is to write a poem. Do not emit introductory text like 'Sure' and other chat. Just write the poem and stop."},
				schema.HumanChatMessage{Content: "write a haiku about the go programming language"},
			}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				return nil
			}))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(completion.Content)
		}
	}
}
