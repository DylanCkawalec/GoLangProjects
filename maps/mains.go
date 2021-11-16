package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":    "ff0000",
		"blue":   "00ff00",
		"yellow": "0000ff",
	}
	printMap(colors)
	fmt.Println(colors)

}

//c is colors
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(" Hex Code for", color, " is ", hex)
	}
}
