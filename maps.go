package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf590",
	}
	colors["white"] = "#f7g890"
	//delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("key is " + key + " and the value is " + value)
	}
}
