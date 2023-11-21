package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
)

var (
	availableModels = []string{
		"falcon",
		"llama2",
		"mistral",
		"mistrallite",
		"orca-mini",
		"zephyr",
	}

	defaultSystemMessage = `Task is to write a poem. Do not emit introductory text like 'Sure' and other chat. Just write the poem and stop.`
	defaultChatMessage   = `write a haiku about the go programming language`

	numSamples    = flag.Int("n", 1, "number of samples to generate")
	systemMessage = flag.String("S", defaultSystemMessage, "system message to use")
	chatMessage   = flag.String("C", defaultChatMessage, "default chat message")
)

type ModelOutput struct {
	Model         string    `json:"model"`
	SystemMessage string    `json:"system"`
	Prompt        string    `json:"prompt"`
	Reply         string    `json:"reply"`
	GeneratedAt   time.Time `json:"t"`
	Elapsed       float64   `json:"elapsed_s"`
}

func main() {
	flag.Parse()
	var outputs []ModelOutput // collect model output
	for _, model := range availableModels {
		llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel(model)))
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		for i := 0; i < *numSamples; i++ {
			started := time.Now()
			completion, err := llm.Call(ctx, []schema.ChatMessage{
				schema.SystemChatMessage{Content: *systemMessage},
				schema.HumanChatMessage{Content: *chatMessage},
			}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				return nil
			}))
			if err != nil {
				log.Fatal(err)
			}
			mo := ModelOutput{
				Model:         model,
				SystemMessage: *systemMessage,
				Prompt:        *chatMessage,
				Reply:         completion.Content,
				GeneratedAt:   time.Now(),
				Elapsed:       time.Since(started).Seconds(),
			}
			outputs = append(outputs, mo)
		}
	}
	enc := json.NewEncoder(os.Stdout)
	for _, mo := range outputs {
		if err := enc.Encode(mo); err != nil {
			log.Fatal(err)
		}
	}
}
