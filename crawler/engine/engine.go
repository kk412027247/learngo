package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

// ...参数类型，表示可以传多个参数进来，并且参数名就是个slice
func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s", r.Url)

		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		// ...是解构slice
		requests = append(requests, parseResult.Requests...)

		// %v 表示打印原本的类型，不转译
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
