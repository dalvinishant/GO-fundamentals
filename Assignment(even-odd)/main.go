package main

import "fmt"

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	num_type := ""

	for _, n := range nums {
		if n%2 == 0 {
			num_type = "even"
		} else {
			num_type = "odd"
		}

		fmt.Printf("%v is %v \n", n, num_type)
	}
}
