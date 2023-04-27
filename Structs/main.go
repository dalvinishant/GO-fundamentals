package main

import "fmt"

// Struct is very similar to dictionary or object which is having some properties/attributes

type person struct {
	firstName string
	lastName  string
	contactinfo
}

// embedding struct
type contactinfo struct {
	email string
	zip   int
}

func main() {
	// struct can be defined in following way

	// With this values in structs are assigned as per their order in struct
	// ninad := person{"Ninad", "K"}

	// ninad := person{
	// 	firstName: "Ninad",
	// 	lastName:  "Kheratkar",
	// }

	// Upon declaration in GO, go assigns variables with a zero-value
	// For e.g : for string it is ''
	//			for int it is 0
	//			for bool it is False
	// var ninad person

	// editing struct
	// ninad.firstName = "Ninad"
	// ninad.lastName = "K"

	// ninad.contactDetails.email = "ninadk@nomail.com"
	// ninad.contactDetails.zip = 411021

	// Declaring embedded structs
	ninad := person{
		firstName: "Ninad",
		lastName:  "K",
		contactinfo: contactinfo{
			email: "ninadk@nomail.com",
			zip:   411021,
		},
	}

	ninad.print()

	// Pass by reference in go

	// create a new variable with address
	// ninadPointer := &ninad

	// call the method with the pointer variable
	// ninadPointer.updateName("Ninu")

	ninad.updateName(("Ninu"))

	ninad.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// GO is referred to as pass-by-value language
// whenever receiver method is called, person p by default referenced as pass by value
// which basically creates a copy of original type
func (pointerPerson *person) updateName(firstName string) {
	(*pointerPerson).firstName = firstName
}
