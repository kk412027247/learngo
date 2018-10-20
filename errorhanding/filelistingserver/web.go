package main

import (
	//"github.com/gpmgo/gopm/modules/log"
	"learngo/errorhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
	// 加了下面这行就可以动态
	_ "net/http/pprof"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 返回函数的参数，跟传入函数的参数一致，就可以直接调用了。
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			// recover 是接住 panic 错误返回信息的地方
			if r := recover(); r != nil {
				log.Printf("Panic: %V", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("error handling request: %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			// http.error 有三个参数，第一个是接受者，第二个是字符串，最后是code本身
			http.Error(writer, http.StatusText(code), code)

		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}

}
