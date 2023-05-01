package main

import "fmt"

// Interface
type bot interface {
	getGreeting() string
}

// Interfaces are not generic types
// Interfaces are implicit in nature i.e. we don't need to do anything explicitly to say xyz method belongs to interface

type englishBot struct{}
type spanishBot struct{}

func main() {
	// Following are concrete types, concrete types could be string, int, etc
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello There!"
}

func (spanishBot) getGreeting() string {

	if a == 0 {
		a = 10
	}

	return "Hola!"
}
