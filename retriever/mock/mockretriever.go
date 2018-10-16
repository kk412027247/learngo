package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

// duck typing 的写法，只要是用了接口的方法，默认就是接口的继承
func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["Contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
