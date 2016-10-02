# scsa 

## What is it?

Semantic web scrapper and analyzer tool.

## Compiling

This software is meant to be a command line / shell binary executable, therefore it has to be compiled first. To compile ``scsa``, you need ``$GOPATH`` to be set and then use tool like [govend](https://github.com/govend/govend) and do ``govend get github.com/spf13/cobra`` before compiling. Once this requirements has been fulfilled, just do:

```bash
go build scsa.go
```

The result should be ``scsa`` binary executable.

## License

```
The MIT License (MIT)
Copyright (c) 2016 WIRG

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
