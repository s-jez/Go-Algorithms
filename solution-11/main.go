//
// Function that takes a number and returns a list of its digits
//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var example3_1 = "123"
	fmt.Println(listOfDigits(example3_1))
	var example3_2 = "400"
	fmt.Println(listOfDigits(example3_2))
}
func listOfDigits(num string) []int {
	var list []int
	for _, i := range num {
		j, err := strconv.Atoi(string(i))
		if err != nil {
			panic(err)
		}
		list = append(list, j)
	}
	return list
}
