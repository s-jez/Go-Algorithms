package main

import "fmt"

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
}

// Sum a list of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
