/*
Package cmoresearch implements a client for C More's search service.

API Documentation

https://cmore-search.b17g.services/docs/

Usage

A small usage example:

		package main

		import (
			"context"
			"fmt"
			"net/url"

			cmoresearch "github.com/TV4/cmoresearch-go"
		)

		func main() {
			client := cmoresearch.NewClient()

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
				if a, ok := hit.(*cmoresearch.Asset); ok {
					fmt.Printf("%s S%02dE%02d\n", a.Brand.TitleSv, a.Season.Number, a.EpisodeNumber)
				}
			}
		}

Output:
		Solsidan S01E01
		Solsidan S01E02
		Solsidan S01E03
*/
package cmoresearch
