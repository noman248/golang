package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "66481811",
		"green": "19848844",
		"white": "4651981",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v: %v\n", color, hex)
	}
}
