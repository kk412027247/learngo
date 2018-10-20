package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) "+
			"Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

}
