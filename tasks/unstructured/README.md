# Unstructured

Sample of 10000 unstructured strings, vary the prompt and model. Bibliographic
records. Try to output JSON.

With curl:

```shell
$ curl -sL 'localhost:11434/api/generate?format=json' -d '{"model": "mistral", "prompt": "Parse the following reference string into JSON: Amis, M. (2001, March 17). A rough trade : The Guardian. Retrieved from The Guardian: http:// www.theguardian.com/books/2001/mar/17/society.martinamis1"}'
```

![](622328.gif)
