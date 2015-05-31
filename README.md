# parse-dockerfile
use native parser of docker builder to parse dockerfile and output json


Unfortunately I'm stuck. Docker engine mixes valition of AST with actual processing the instruction within one function. Therefore I cannot incorporate validation in this project.

The way docker engine parses dockerfile is that

1. it parses dockerfile into AST
2. it takes the AST and for every instruction it validates the input and process' it

Unfortunately, this is pretty bad, because the validation cannot be used within this project, because it is part of a instruction dispatcher within engine. It would much better to validate input outside of dispatcher.


# Usage

```
% go build && ./parse-dockerfile ./test/Dockerfile
```

Get docker's parser

```
% go get github.com/docker/docker/builder/parser
```
