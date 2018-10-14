package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface{
	Get(url string) string
}


func download (r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}

	inspect(r)

	// Type assertion 可以通过.()的方式，取到具体名字的值
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever)
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("can not ge the assertion")
	}



	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
