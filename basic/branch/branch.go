package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Worngscore: :%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score <= 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g

}

func main() {
	const filename = "basic/branch/bcd.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Printf(
		grade(0),
		grade(59),
		grade(100),
	)
}