package main

import "fmt"

func main() {
	fmt.Println(NumInList([]int{5, 3, 2, 1}, 5))
}

// Find a number in a list of numbers
func NumInList(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
