package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Function declaration needs to have return type
func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// Receiver on a function -> Any variable of type "deck" now gets access to "print" method defined below

// Receiver format
// func (arg type) method_name() {} -> Similar to self, this
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Type Conversion in GO is similar to that of Python
// Format := data_type_to_convert_to(value_we_have_currently)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	// Following 2 lines of code is specific to "Rand" Standard Package in GO
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i, _ := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}

func newDeckFromFile(filename string) deck {

	// ReadFile returns []byte, error as return values
	bs, err := ioutil.ReadFile(filename)

	// nil is value in GO that is similar to 'None'(from Python)
	if err != nil {

		// Log Error
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	// type conversion is possible for a custom type here is because, 'deck' is just an extension to []string
	return deck(strings.Split(string(bs), ","))
}
