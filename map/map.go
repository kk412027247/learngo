package main

import "fmt"

func main() {
	dict := make(map[string]int)

	dict2 := map[string]string{"red":"#da1337", "orange":"#e95a22"}

	fmt.Println(dict, dict2)

	colors := map[string]string{}
	colors["red"] = "#da1337"

	fmt.Println(colors)

	colors1 := map[string]string{
		"aliceBlue":"#f0f8ff",
		"Coral":"#ff7f50",
		"DarkGray":"#a9a9a9",
		"ForestGreen":"#228b22",
	}

	for key,value := range colors1 {
		fmt.Printf("key: %s value: %s\n", key, value)
	}
	fmt.Println()
	removeColor(colors1, "Coral")

	for key,value := range colors1 {
		fmt.Printf("key: %s value: %s\n", key, value)
	}

}

func removeColor (colors map[string]string, key string) {
	delete(colors, key)
}
