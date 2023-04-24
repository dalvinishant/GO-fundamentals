package main

import "fmt"

func main() {
	// DECLARATION FORMAT
	//		var				some_var			type										= some_value
	//	new variable		variable_name		indication to complier						Value assigned to the variable
	//											that only given string would be assigned
	// Some Basic fundamental types available in GO include
	// 1. bool [ Boolean ]
	// 2. int [ Integer ]
	// 3. string [ String ]
	// 4. float64 [ Float ]
	// 5. map [ Hash Map ]

	// var card string = "Ace of Spades"
	// Above line can also be rewritten as
	// card := "Ace of Spades"

	// Slices in GO
	// cards := newDeck() // -> Symbol `:=` to be used only if we are creating ans assigning the value to a new variable
	// card = "Five of Diamonds" -> Symbol `=` for reassignment

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")

	cards.print()

	fmt.Println("\n After Shuffle")

	cards.shuffle()
	cards.print()

	// Unpacking the value returned by function deal
	// hand, remaingingDeck := deal(cards, 5)

	// // LOOP FORMAT
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// Calling a receiver for type "deck"

	// hand.print()
	// remaingingDeck.print()

	// fmt.Println(cards)
}

// **** Notes ****
// Since GO is a statically typed languaged, hence it requires us to explicitly mention the types
// List(python) in GO is available in 2 options
// 1. Array
//		Fixed length
// 2. Sclies
//		Array that can grow or shrink
//		Each object in slice, must be of same type
//		Slices are zero-indexed (i.e. their indexing starts from 0)
// GO is not an Object Orient language
