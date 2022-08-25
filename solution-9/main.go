//
// Program that returns the odd elements in a list
//

package main

import "fmt"

func main() {
	var example1_1 = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(oddPosition(example1_1))
	var example1_2 = []int{1, -1, 2, -2}
	fmt.Println(oddPosition(example1_2))
}
func oddPosition(arr []int) []int {
	var output []int
	for _, num := range arr {
		if num%2 != 0 {
			output = append(output, num)
		}
	}
	fmt.Print("output ")
	return output
}
