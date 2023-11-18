# Code structure

* LLM is an [interface](https://github.com/jmorganca/ollama/blob/984714f13182a6b32a8d456427536bbe51d6032e/llm/llm.go#L16-L24)

```
type LLM interface {
	Predict(context.Context, []int, string, string, func(api.GenerateResponse)) error
	Embedding(context.Context, string) ([]float64, error)
	Encode(context.Context, string) ([]int, error)
	Decode(context.Context, []int) (string, error)
	SetOptions(api.Options)
	Close()
	Ping(context.Context) error
}
```

* DecodeGGML from a ReadSeeker: [ggml.go:182](https://github.com/jmorganca/ollama/blob/984714f13182a6b32a8d456427536bbe51d6032e/llm/ggml.go#L182-L212)
* GGML is just a slim struct: [ggml.go:9](https://github.com/jmorganca/ollama/blob/984714f13182a6b32a8d456427536bbe51d6032e/llm/ggml.go#L9-L13)

```go
type GGML struct {
    magic uint32
    container
    model
}
```

Container is an even slimmer interface:

```go
type container interface {
    Name() string
    Decode(io.Reader) (model, error)
}
```

Model is only metadata:

```go
type model interface {
    ModelFamily() string
    ModelType() string
    FileType() string
    NumLayers() int64
}
```

Format names:

* GGML (low level ML library)
* GGMF (?)
* GGJT (via: [llama.cpp](https://github.com/ggerganov/llama.cpp), [#613](https://github.com/ggerganov/llama.cpp/pull/613))

In [convert-llama-ggml-to-gguf.py](https://github.com/ggerganov/llama.cpp/blob/2923f17f6fec049a71186636c3c4d96408856194/convert-llama-ggml-to-gguf.py#L17-L20):

```python
class GGMLFormat(IntEnum):
    GGML = 0
    GGMF = 1
    GGJT = 2
```
