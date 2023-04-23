// Command to run/compile GO files
// 1. go run file_name.go -> compiles + executes the file
// 2. go build file_name.go -> compiles the given file
//		- Post Compilation, running `./file_name` would actually execute the file

// PACKAGE DECLARATION SECTION
// =================================
// name of the package this file belongs to
package main

// There 2 types of packages in GO
// 1. Executable -> Generates a file that we can run
//		- To create a executable package we explicitly use the keyword `main` in package declaration
//		- `main` name is reserved/sacred to identify if a file is a execute package or resuable
//		- When an executable package is given to build, an executable file is generated (this only happens if main is used)
// 2. Resuable -> Code used as 'helpers'. Good place to put resuable logic
//		- Compiling a resuable package does not generate a executible
//===================================

// IMPORT SECTION
// ==================================
// Standard package in GO
import "fmt" // Shortened form of 'format'
// Checkout : https://golang.org/pkg to know more about packages
// ==================================

// FUNCTIONS SECTION
// ==================================
func main() {
	fmt.Println("Hello World!")
}

// ==================================
