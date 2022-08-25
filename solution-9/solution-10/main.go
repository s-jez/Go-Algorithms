//
// Program that returns the cumulative sum of elements in the list
//
package main

import "fmt"

func main() {
	var example2_1 = []int{1, 1, 1}
	fmt.Println(cumulativeSum(example2_1))
	var example2_2 = []int{1, -1, 3}
	fmt.Println(cumulativeSum(example2_2))

}
func cumulativeSum(arr []int) []int {
	var total int
	output := []int{}
	for x := range arr {
		total += arr[x]
		output = append(output, total)
	}
	fmt.Print("output ")
	return output
}
