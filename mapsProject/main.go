package main

import "fmt"

func main() {
	//maps syntax map[keyType]valType
	// colors := map[string]string{
	// "red":   "#ff0000",
	// 	// "green": "#4bf745",
	// }

	//make (like new)
	colors := make(map[string]string)
	//crear key:value pair
	colors["white"] = "#ffffff"
	colors["black"] = "#00000"
	// delete some map key:vale pair
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Color name: %v , Color hex: %v\n", key, value)
	}
}
