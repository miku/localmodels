# Running local models

> 2023-11-21, Martin Czygan <martin.czygan@gmail.com>, [L Gopher](https://golangleipzig.space), Open Data Engineer at [IA](https://archive.org)

Short talk about running local models, using Go tools.

## Personal Timeline

> "What a difference a week makes"

* on 2022-11-30, [chatGPT](https://en.wikipedia.org/wiki/ChatGPT) is released
* on 2022-12-12 (+2w), one week after tweet ID [1599971348717051904](https://twitter.com/alexandr_wang/status/1599971348717051904), we discuss the new role of "prompt engineer" in a CS class at [LU Leipzig](https://en.wikipedia.org/wiki/Lancaster_University_Leipzig)

> I am going to assert that Riley is the first Staff Prompt Engineer hired *anywhere*.

* on 2023-02-14 (+9w), I ask a question on how long before we can run things locally at the [Leipzig Python User Group](https://lpug.github.io/) -- personally, I expected 1-3 years timeline
* on 2023-04-18 (+9w), we discuss C/GO and ggml (ai-on-the-edge) at [Leipzig Gophers #35](https://golangleipzig.space/posts/meetup-35-wrapup/)
* on 2023-07-20 (+13w), ollama is released (two models), [HN](https://news.ycombinator.com/item?id=36802582)
* on 2023-11-21 (+17w), today, 100s of models (TODO: understand GGUF spread)

## 73 Anno TT

TODO: ML/DL is conceptually a black box, with simple parts and complex emergent
behavious; it has been compared to thermodynamics, developed out of the steam
engines with simple particles and laws leadning to complex, statistical,
aggregate behaviour.

I love technology, still not ❤️  black boxes - in Year 73 after [Computing Machinery
and Intelligence](https://phil415.pbworks.com/f/TuringComputing.pdf). Anno 73 TT.

From [Nature](https://www.nature.com/), 2023-07-23: [Understanding ChatGPT is a bold new challenge for science](https://www.nature.com/articles/d41586-023-02366-2.pdf)

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

Go is a nice infra language, what projects exist for model infra?

* going to look at a tool, from the *outside* and a bit from the *inside*

## POLL

* have you ran a local LLM, yes or no?
* if so, which models?

## OLLAMA

* first appeared in [07/2023](http://web.archive.org/web/20230720133902/https://ollama.ai/) (~18 weeks ago)
* very inspired by docker, not images, but models
* built on [llama](https://ai.meta.com/llama/) (meta), [GGML](http://ggml.ai/) ai-on-the-edge ecosystem, especially using [GGUF](https://www.reddit.com/r/LocalLLaMA/comments/15triq2/gguf_is_going_to_make_llamacpp_much_better_and/) - a unified image format
* docker may be considered less a glorified [nsenter](https://man7.org/linux/man-pages/man1/nsenter.1.html), but more (lots of) glue to go from spec to image to process, code lifecycle management
* very clean user experience (that many projects lack)

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


### TODO

* [ ] define a couple of tasks and run them in batch against the API
* [ ] create custom prompts, example tasks

Some specific prompts may be:

* [ ] an instructor for a specific programming style (e.g. see elements of programming style)


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

* `/api/generate/`


## Credits

* [Richard Feynman's blackboard at time of his death](https://digital.archives.caltech.edu/collections/Photographs/1.10-29/) (1988)
