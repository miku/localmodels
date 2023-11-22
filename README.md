# Testdriving OLLAMA

> 2023-11-21, Leipzig Gophers [Meetup #38](#), [Martin Czygan](https://www.linkedin.com/in/martin-czygan-58348842/), [L Gopher](https://golangleipzig.space), Open Data Engineer at [IA](https://archive.org)

Short talk about running local models, using Go tools.

## Personal Timeline

> "What a difference a week makes"

* on 2022-11-30, [chatGPT](https://en.wikipedia.org/wiki/ChatGPT) is released
* on 2022-12-12 (+2w), one week after tweet ID [1599971348717051904](https://twitter.com/alexandr_wang/status/1599971348717051904), we discuss the new role of "prompt engineer" in a CS class at [LU Leipzig](https://en.wikipedia.org/wiki/Lancaster_University_Leipzig)

> I am going to assert that Riley is the first Staff Prompt Engineer hired *anywhere*.

* on 2023-02-14 (+9w), I ask a question on how long before we can run things locally at the [Leipzig Python User Group](https://lpug.github.io/) -- personally, I expected 2-5 years timeline
* on 2023-04-18 (+9w), we discuss C/GO and ggml (ai-on-the-edge) at [Leipzig Gophers #35](https://golangleipzig.space/posts/meetup-35-wrapup/)
* on 2023-07-20 (+13w), [ollama](https://ollama.ai) is released (with two models), [HN](https://news.ycombinator.com/item?id=36802582)
* on 2023-11-21 (+17w), today, 43 models (each with a couple of tags/versions)

## Confusion

Turing Test was proposed in 1950. From [Nature](https://www.nature.com/),
2023-07-23: [Understanding ChatGPT is a bold new challenge for
science](https://www.nature.com/articles/d41586-023-02366-2.pdf)

> This **lack of robustness** signals a lack of **reliability** in the real world.

![](static/default-30.jpg)

> What I cannot create, I do not understand.

Open models not binary:

> We propose a framework to assess six levels of access to generative AI
> systems, from [The Gradient of Generative AI Release: Methods and
> Considerations](https://arxiv.org/pdf/2302.04844.pdf):

* fully closed
* gradual or staged access
* hosted access
* cloud-based or API access
* downloadable access and
* fully open.

A [prolific AI
Researcher](https://scholar.google.de/citations?user=x04W_mMAAAAJ&hl=en) (with
387K citations in the past 5 years) believes open source AI is ok for less
capable models: [Open-Source vs. Closed-Source
AI](https://www.youtube.com/watch?v=ZfYrJlfLs1Q)

For today, let's focus on Go. Go is a nice infra language, what projects exist
for model infra?

* going to look at a tool, from the *outside* and a bit from the *inside*

## POLL

* have you written a [markov chain based text generator](https://go.dev/doc/codewalk/markov/) in Go?
* have you ran a local LLM, yes or no?
    * (only) about 10% said yes
* if so, any particular model or tool?

## OLLAMA

* first appeared in [07/2023](http://web.archive.org/web/20230720133902/https://ollama.ai/) (~18 weeks ago)
* very inspired by docker, not images, but models
* built on [llama](https://ai.meta.com/llama/) (meta), [GGML](http://ggml.ai/) ai-on-the-edge ecosystem, especially using [GGUF](https://www.reddit.com/r/LocalLLaMA/comments/15triq2/gguf_is_going_to_make_llamacpp_much_better_and/) - a unified image format
* docker may be considered less a glorified [nsenter](https://man7.org/linux/man-pages/man1/nsenter.1.html), but more (lots of) glue to go from spec to image to process, code lifecycle management; similarly ollama may be a way to organize the ai "model lifecycle"
* clean developer UX

### Time-to-chat

From zero to chat in about 5 minutes, on a [power-efficient
CPU](https://www.intel.com/content/www/us/en/processors/processor-numbers.html).
Started w/ 2 models, as of 11/2023 hosting 36+ models, a docker-like model.

```
$ git clone git@github.com:jmorganca/ollama.git
$ cd ollama
$ go generate ./... && go build . # cp ollama ...
```

Follows a client server model, like docker.

```
$ ollama serve
```

Once it is running, we can pull models.

```
$ ollama pull llama2
pulling manifest
pulling 22f7f8ef5f4c... 100% |..
pulling 8c17c2ebb0ea... 100% |..
pulling 7c23fb36d801... 100% |..
pulling 2e0493f67d0c... 100% |..
pulling 2759286baa87... 100% |..
pulling 5407e3188df9... 100% |..
verifying sha256 digest
writing manifest
removing any unused layers
success
```

### Some examples

```
$ ollama run zephyr
>>> please complete: {"author": "Turing, Alan", "title" ... }

{
  "author": "Alan Turing",
  "title": "On Computable Numbers, With an Application to the Entscheidungsproblem",
  "publication_date": "1936-07-15",
  "journal": "Proceedings of the London Mathematical Society. Series 2",
  "volume": "42",
  "pages": "230–265"
}
```

Formatting mine.

### More

The whole [prompt engineering](https://en.wikipedia.org/wiki/Prompt_engineering) thing is kind of mysterious to me. Do you get better output by showing emotions?

* [Large Language Models Understand and Can Be Enhanced by Emotional Stimuli](https://arxiv.org/pdf/2307.11760.pdf) -- "EmotionPrompt"

> To this end, we first conduct automatic experiments on 45 tasks using various
> LLMs, including Flan-T5-Large, Vicuna, Llama 2, BLOOM, ChatGPT, and GPT-4.

### Batch Mode

```
[GIN-debug] POST   /api/pull       --> gith...m/jmo...ma/server.PullModelHandler (5 handlers)
[GIN-debug] POST   /api/generate   --> gith...m/jmo...ma/server.GenerateHandler (5 handlers)
[GIN-debug] POST   /api/embeddings --> gith...m/jmo...ma/server.EmbeddingHandler (5 handlers)
[GIN-debug] POST   /api/create     --> gith...m/jmo...ma/server.CreateModelHandler (5 handlers)
[GIN-debug] POST   /api/push       --> gith...m/jmo...ma/server.PushModelHandler (5 handlers)
[GIN-debug] POST   /api/copy       --> gith...m/jmo...ma/server.CopyModelHandler (5 handlers)
[GIN-debug] DELETE /api/delete     --> gith...m/jmo...ma/server.DeleteModelHandler (5 handlers)
[GIN-debug] POST   /api/show       --> gith...m/jmo...ma/server.ShowModelHandler (5 handlers)
[GIN-debug] GET    /               --> gith...m/jmo...ma/server.Serve.func2 (5 handlers)
[GIN-debug] GET    /api/tags       --> gith...m/jmo...ma/server.ListModelsHandler (5 handlers)
[GIN-debug] HEAD   /               --> gith...m/jmo...ma/server.Serve.func2 (5 handlers)
[GIN-debug] HEAD   /api/tags       --> gith...m/jmo...ma/server.ListModelsHandler (5 handlers)
```

Specifically `/api/generate/`

### Constraints

* possible to enforce JSON generation

### Customizing models

> weights, configuration, and data in a single package

Using a Modelfile.

```
FROM llama2
# sets the temperature to 1 [higher is more creative, lower is more coherent]
PARAMETER temperature 1
# sets the context window size to 4096, this controls how many tokens the LLM can use as context to generate the next token
PARAMETER num_ctx 4096

# sets a custom system prompt to specify the behavior of the chat assistant
SYSTEM You are Mario from super mario bros, acting as an assistant.
```

Freeze this as a custom package:

```shell
$ ollama create llama-mario -f custom/Modelfile.mario
$ ollama run llama-mario
```

About 16 parameters to tweak: [Valid Parameters and Values](https://github.com/jmorganca/ollama/blob/main/docs/modelfile.md#valid-parameters-and-values)

## Task 1: "haiku"

* generate a small volume of Go programming haiku

```
// haikugen generates
// JSON output for later eval
// cannot parallelize
```

* [haikugen.go](https://github.com/miku/localmodels/blob/main/tasks/haiku/haikugen.go)

## Task 2 "bibliography"

* given unstructured strings, parse the to json
* [unstructured](https://github.com/miku/localmodels/tree/main/tasks/unstructured)

## Credits

* [Richard Feynman's blackboard at time of his death](https://digital.archives.caltech.edu/collections/Photographs/1.10-29/) (1988)
