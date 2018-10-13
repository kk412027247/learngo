package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)
	fmt.Println("Traversiong map")
	for k,v := range m {
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if caurseName, ok := m["caurse"]; ok{
		fmt.Println(caurseName, ok)
	} else {
		fmt.Println("key cause does not exit")
	}


	fmt.Println("Deleting values")
	name,ok := m["name"]
	println(name, ok)
	delete(m, "name")
	name,ok = m["name"]
	println(name, ok)

}
