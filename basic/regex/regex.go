package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse1@gaimal.com@D.COM
My email is 322@gaimal.com@D.COM
My email is 333@222.com@D.COM
My email is 333@222.com.cn@D.COM.CN
`

func main() {
	re := regexp.MustCompile(`([\w]+)@(\w+)(\.[\w.]+)`)
	s := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(s)

	for _, m := range s {
		fmt.Println(m)
	}

}
