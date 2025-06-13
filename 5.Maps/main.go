package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// 	"white": "#ffffff",
	// }
	// var colors map[string]string
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	// delete(colors, "white")
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Color:", key, "Hex:", value)
	}
}
