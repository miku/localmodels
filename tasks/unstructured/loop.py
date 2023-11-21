#!/usr/bin/env python

# localhost:11434/api/generate?format=json' -d '{"model": "mistral", "prompt":
# "Parse the following reference string into JSON: Amis, M. (2001, March
# 17). A rough trade : The Guardian. Retrieved from The Guardian: http://
# www.theguardian.com/books/2001/mar/17/society.martinamis1"}

# 33s


import argparse
import collections
import fileinput
import io
import json
import requests
import sys

DEFAULT_OLLAMA_URL = "http://localhost:11434/api/generate"

if __name__ == '__main__':

    parser = argparse.ArgumentParser()
    parser.add_argument("--api", default=DEFAULT_OLLAMA_URL, help="ollama generate api url")
    parser.add_argument("--model", default="mistral", help="model name")
    parser.add_argument('--prefix', default="Parse the following reference string into JSON: ", help="prefix for prompt")
    parser.add_argument('files', metavar='FILE', nargs='*', help='files to read, if empty, stdin is used')

    args = parser.parse_args()

    # example line
    # "Parse the following reference string into JSON: Amis, M. (2001, March 17). A rough trade : The Guardian. Retrieved from The Guardian: http:// www.theguardian.com/books/2001/mar/17/society.martinamis1"

    for line in fileinput.input(files=args.files if len(args.files) > 0 else ('-', )):
        if not line.strip():
            continue
        stats = collections.Counter()
        params = {
            "format": "json",
            "model": args.model,
            "prompt": args.prefix + line,
        }
        sio = io.StringIO()
        # generate
        r = requests.post(args.api, json=params, stream=True)
        for line in r.iter_lines():
            if line:
                resp = json.loads(line)
                print(resp.get("response"), end="", file=sio)
                if resp.get("done", False):
                    break
        # validate
        generated = sio.getvalue()
        try:
            v = json.loads(generated)
            print(json.dumps({
                "input": line,
                "parsed": v,
            }))
        except Exception as exc:
            stats["failed"] += 1
        else:
            stats["ok"] += 1
    # print(json.dumps(stats), file=sys.stderr)
