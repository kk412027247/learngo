package main

import "fmt"

var (
	aa = 3
	ss = "kkk"
	bb = true
)


func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a, s)
}

func variableInitialValue(){
	var a = 3
	var s = "abc"
	fmt.Println(a, s)
}

func variableShorter(){
	a,b,c,s := 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
	fmt.Println(aa, ss, bb)
}
