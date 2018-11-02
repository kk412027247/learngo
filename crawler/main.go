package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/paser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: paser.ParseCityList,
	})
}
