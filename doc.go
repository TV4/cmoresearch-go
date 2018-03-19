/*
Package search implements a client for Bonnier Broadcasting's search service.

API Documentation

https://cmore-search.b17g.services/docs/

Usage

A small usage example:

		package main

		import (
			"context"
			"fmt"
			"net/url"

			search "github.com/TV4/search-go"
		)

		func main() {
			client := search.NewClient()

			res, err := client.Search(
				context.Background(),
				url.Values{
					"device_type": {"tve_web"},
					"lang":        {"sv"},
					"site":        {"cmore.se"},

					"brand_id":  {"34515"},
					"season":    {"1"},
					"sort_by":   {"episode_number"},
					"order":     {"asc"},
					"page_size": {"3"},
				},
			)

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, hit := range res.Hits {
				if a, ok := hit.(*search.Asset); ok {
					fmt.Printf("%s S%02dE%02d\n", a.Brand.TitleSv, a.Season.Number, a.EpisodeNumber)
				}
			}
		}

Output:
		Solsidan S01E01
		Solsidan S01E02
		Solsidan S01E03
*/
package search
