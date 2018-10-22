package main

import (
	"learngo/rss/search"
	"log"
	"os"
	_ "learngo/rss/matchers"
)

func init(){
	log.SetOutput(os.Stdout)
}

func main() {
  	search.Run("president")
}
