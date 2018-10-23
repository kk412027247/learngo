package search

import "log"

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string)([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result){
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println("搜索出错了",err)
		return
	}

	for _, result := range searchResult {
		results <- result
	}
}

func Display(results chan *Result){
	for result := range results {
		log.Printf("搜索结果 %s:\n%s\n\n",result.Field, result.Content)
	}
}
