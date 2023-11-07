# Running local models

> 2023-11-21, Martin Czygan <martin.czygan@gmail.com>, [L Gopher](https://golangleipzig.space), Open Data Engineer at [IA](https://archive.org)

Short talk about running local models, using Go tools.

## Motivation

I love technology, still not ❤️  black boxes - in Year 73 after [Computing Machinery
and Intelligence](https://phil415.pbworks.com/f/TuringComputing.pdf).

From 2023-07-23: [Understanding ChatGPT is a bold new challenge for science](https://www.nature.com/articles/d41586-023-02366-2.pdf)

> This **lack of robustness** signals a lack of **reliability** in the real world.

![](static/default.jpg)

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

## POLL

* have you ran a local LLM, yes or no?
* if so, which models?

## OLLAMA

From zero to chat in about 5 minutes. As of 11/2023 hosting 36+ models, a
docker-like model.

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
$ ollama pull llama2                                                                                                                                                                                                         [33/33]
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




## Credits

* [Richard Feynman's blackboard at time of his death](https://digital.archives.caltech.edu/collections/Photographs/1.10-29/) (1988)
