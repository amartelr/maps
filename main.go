package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	//delete(colors, "red")
	printMap(colors)

	anothercolor := make(map[string]string)
	anothercolor["white"] = "#fffff"
	fmt.Println(anothercolor)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
