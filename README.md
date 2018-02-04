# Wham !

[![Build Status](https://travis-ci.org/juanwolf/wham.svg?branch=master)](https://travis-ci.org/juanwolf/wham) [![codecov](https://codecov.io/gh/juanwolf/wham/branch/master/graph/badge.svg)](https://codecov.io/gh/juanwolf/wham)

Wake me up before you go go :notes:. Wham is a little cli to track your overtime when your on call.
Not easy to track that down at 4am in the morning, right?

## Installation

You can donwload the binaries from github directly or install it manually.

```
go get -u github.com/golang/dep/cmd/dep
go get github.com/juanwolf/wham
cd $GOPATH/github.com/juanwolf/wham
dep ensure
go install
```

## Usage

First, you need to start wham when you receive a call:

```
wham start
```

Once fixed, you can stop wham:

```
wham stop
```

And you should have a lovely output telling you how long you worked for.

Ok this is not crazy _yet_ but will be in the future releases!


## Plan

1. Save all the sessions in a file every month
2. Export functions to generate xml, spreadsheets, or even call an api
3. Add flags to specify why you got up, like `wham start -m "One DC is burning"`
4. Add configuration in a  ~/.wham



## License

MIT License
