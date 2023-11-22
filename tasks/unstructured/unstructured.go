// haikugen generates
// JSON output for later eval
// cannot parallelize
//
// data points (Xeon, PM: 2338/35224)
//
// 6 models x 10 message each: 1m53.604s, generation took between 8 and 0.04
// seconds, output between 849 and 2 chars (14, or 25%). Around 3000% CPU usage.
//
// 6 models x 200 messages: 1200 replies in 34min or 0.58 haikus per second.
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

	defaultSystemMessage = `Parse reference strings into JSON`
	defaultChatMessage   = `Amis, M. (2001, March 17). A rough trade : The Guardian. Retrieved from The Guardian: http:// www.theguardian.com/books/2001/mar/17/society.martinamis1`

	timeout       = flag.Duration("t", 30*time.Second, "timeout")
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
	enc := json.NewEncoder(os.Stdout)
	for _, model := range availableModels {
		llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel(model)))
		if err != nil {
			log.Fatal(err)
		}
		success := 0
		for i := 0; i < *numSamples; i++ {
			ctx, cancelFunc := context.WithTimeout(context.Background(), *timeout)
			defer cancelFunc()
			started := time.Now()
			completion, err := llm.Call(ctx, []schema.ChatMessage{
				schema.SystemChatMessage{Content: *systemMessage},
				schema.HumanChatMessage{Content: *chatMessage},
			}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				return nil
			}))
			if err != nil {
				log.Printf("%s failed: %v, skipping", model, err)
				continue
			}
			success++
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
		log.Printf("%d/%d succeeded for %s", success, *numSamples, model)
	}
}
