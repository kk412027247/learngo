package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	// 最后面的string表示返回值是string
	Post(url string, from map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {

	s.Post(url, map[string]string{
		"Contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	// 这里的r之所以能打印出字符串，是因为Retriever 实现了string方法
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	// 这里的 mock retriever 这个struct 同时实现了  Retriever 与   Poster 接口的方法，所以根据duck typing 原理
	// 这个  retriever 就是 同时继承了 Retriever 与 Poster
	// r 只是 Retriever的继承，无法同时调用 get 和 post方法，但是可以
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	// Type assertion 可以通过.()的方式，取到具体名字的值
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever)
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("can not ge the assertion")
	}

	fmt.Println(download(r))
	fmt.Println("try a session")

	fmt.Println(session(&retriever))
}
