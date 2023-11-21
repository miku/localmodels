// haikugen generates
// JSON output for later eval
// cannot parallelize
//
// data points:
//
// 6 models x 10 message each: 1m53.604s, generation took between 8 and 0.04
// seconds, output between 849 and 2 chars (14, or 25%). Around 3000% CPU usage.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
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

	timeout       = flag.Duration("t", 8*time.Second, "timeout")
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
	totalRuns := len(availableModels) * *numSamples
	bar := progressbar.NewOptions(totalRuns, progressbar.OptionSetWriter(os.Stderr))
	enc := json.NewEncoder(os.Stdout)
	for _, model := range availableModels {
		llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel(model)))
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancelFunc := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancelFunc()
		for i := 0; i < *numSamples; i++ {
			bar.Add(1)
			started := time.Now()
			completion, err := llm.Call(ctx, []schema.ChatMessage{
				schema.SystemChatMessage{Content: *systemMessage},
				schema.HumanChatMessage{Content: *chatMessage},
			}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				return nil
			}))
			if err != nil {
				log.Printf("%d failed: %v, skipping", model, err)
				continue
			}
			mo := ModelOutput{
				Model:         model,
				SystemMessage: *systemMessage,
				Prompt:        *chatMessage,
				Reply:         completion.Content,
				GeneratedAt:   time.Now(),
				Elapsed:       time.Since(started).Seconds(),
			}
			if err := enc.Encode(mo); err != nil {
				log.Fatal(err)
			}
		}
	}
}
