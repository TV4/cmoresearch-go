# cmoresearch-go

[![Build Status](https://travis-ci.org/TV4/cmoresearch-go.svg?branch=master)](https://travis-ci.org/TV4/cmoresearch-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/cmoresearch-go)](https://goreportcard.com/report/github.com/TV4/cmoresearch-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/cmoresearch-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/cmoresearch-go#license)

`cmoresearch-go` is a Go client for C More's search service.

## Installation
```
go get -u github.com/TV4/cmoresearch-go
```

## Usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	cmoresearch "github.com/TV4/cmoresearch-go"
)

func main() {
	client := cmoresearch.NewClient(
		cmoresearch.SetDebugLogf(log.New(os.Stderr, "", 0).Printf),
	)

	query := url.Values{
		"site":        {"cmore.se"},
		"video_ids":   {"2222333,2222334"},
	}

	res, err := client.Search(context.Background(), query)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("%d hits\n", res.TotalHits)

	for _, hit := range res.Hits {
		switch h := hit.(type) {
		case *cmoresearch.Asset:
			fmt.Printf("Asset %s\n", h.VideoID)
		case *cmoresearch.Series:
			fmt.Printf("Series %s\n", h.BrandID)
		}
	}
}
```

```
GET https://cmore-search.b17g.services/search?site=cmore.se&video_ids=2222333%2C2222334
2 hits
Asset 2222334
Asset 2222333
```

## API Documentation

https://cmore-search.b17g.services/docs/

## License

Copyright (c) 2017-2019 TV4

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
