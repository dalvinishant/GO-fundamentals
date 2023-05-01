package main

import "fmt"

type custom_map map[string]string

func main() {

	// Maps in GO are similar to dict in python
	// Only difference are the following 2 points
	//		1. All the keys must be of the same type
	//		2. All the values must be of the same type
	//		3. Keys and values can be of different types

	// syntax for declaring map

	// variable_name := map[key_type] value_type{}
	// For e.g.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf545",
		"white": "#fffff",
	}

	// var variable_name map[key_type]value_type
	// For e.g.
	// var colors map[string]string

	// variable_name := make(map[key_type] value_type)
	// colors := make(map[string]string)

	// Assigning value in map
	colors["white"] = "#ffffff"

	//difference between struct and map is that, maps uses "." notaion for access keys in struc
	// whereas, map uses "[]" to access its keys

	// Deleting a key in map
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v color is %v \n", color, hex)
	}
}
