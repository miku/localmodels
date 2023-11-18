# Models

Where to all the models come from?

* original LLAMA paper: [https://arxiv.org/pdf/2302.13971.pdf](https://arxiv.org/pdf/2302.13971.pdf), 2023-02-27

> We introduce LLaMA, a collection of foundation language models ranging from 7B
to 65B parameters. We train our models on trillions of tokens, and show that it
is possible to train state-of-the-art models using publicly avail- able
datasets exclusively, without resorting to proprietary and inaccessible
datasets. In particular, LLaMA-13B outperforms GPT-3 (175B) on most benchmarks,
and LLaMA- 65B is competitive with the best models, Chinchilla-70B and
PaLM-540B. We release all our models to the research community.

Tradeoff between training and inference time.

For instance, although Hoffmann et al. (2022) recommends training a 10B model
on 200B tokens, we find that the performance of a 7B model continues to improve
even after 1T tokens.

> How much are 1T tokens?

* Assume 1 token is 1 word: 1T words. EN wikipedia:
[4.3B](https://en.wikipedia.org/wiki/Wikipedia:Size_of_Wikipedia) words. So
about 233x english wikipedias.
* Average novel has 80000 words, so 12.5M novels. DNB has [17.5M](https://www.dnb.de/DE/Sammlungen/Buecher/buecher_node.html) books.
* Or software: [Software Heritage Archive](https://archive.softwareheritage.org/)
has 17,276,493,075 files, assume some average SLOC of 200 and 5 tokens per line, we would have: 17,276,493,075,000 tokens (17T)

Smaller model, trained longer:

> for instance, LLaMA-13B outperforms GPT-3 on most benchmarks, despite being 10x smaller.

Interestingly, the raw data used for training is not that large, probably not
exceeding 5TB (compressed).

----

One dataset is Stack Exchange; let's just use SO, which is the largest:

```
my contributions:                 1378 Q/A
total questions:              24000000 [W](https://en.wikipedia.org/wiki/Stack_Overflow)
total answers:                35000000 [W](https://en.wikipedia.org/wiki/Stack_Overflow)
total contributions           59000000
my contribution ratio:               0.0000276
LLAMA sampling ratio from SE:        0.2 [A](https://arxiv.org/pdf/2302.13971.pdf)
my contribution to llama:            0.000005520 or 0.000552%
```

----

LLAMA training notes:

> When training a 65B-parameter model, our code processes around 380
tokens/sec/GPU on 2048 A100 GPU with 80GB of RAM. This means that training over
our dataset containing 1.4T tokens takes approximately 21 days.

Current (2023) price of A100: EUR 18587 (["Nur noch 2 auf
Lager"](https://is.gd/174ewV)) - so the gear may cost up to: EUR38M. Renting
2048 GPUs A100 40GB for a day costs [78643](https://puzl.cloud/gpu-cloud).
Usage for 21 days would be: $1.6M or $3.3M.

Assuming we see some kind of Moore's Law play out; translate to "halving the
price" every two years, in 2030 the gear could cost less than 5M and the
training cost could be at $200-400K.
