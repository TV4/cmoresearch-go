# Search

[![Build Status](https://travis-ci.org/TV4/search-go.svg?branch=master)](https://travis-ci.org/TV4/search-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/search-go)](https://goreportcard.com/report/github.com/TV4/search-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/search-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/search-go#license)

Search is a Go client for Bonnier Broadcasting's search service.

## Installation
```
go get -u github.com/TV4/search-go
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

	search "github.com/TV4/search-go"
)

func main() {
	client := search.NewClient(
		search.SetDebugLogf(log.New(os.Stderr, "", 0).Printf),
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
		case *search.Asset:
			fmt.Printf("Asset %s\n", h.VideoID)
		case *search.Series:
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
