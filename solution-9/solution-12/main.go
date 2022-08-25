//
// Return the centered average of an array of ints, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and likewise for the largest value.
//

package main

import "fmt"

func main() {
	var example4_1 = []int{1, 2, 3, 4, 100}
	fmt.Println(centeredAvg(example4_1))
	var example4_2 = []int{1, 1, 5, 5, 10, 8, 7}
	fmt.Println(centeredAvg(example4_2))
	var example4_3 = []int{-10, -4, -2, -4, -2, 0}
	fmt.Println(centeredAvg(example4_3))
}
func centeredAvg(arr []int) int {
	var sum int
	var min = arr[0]
	var max = arr[0]

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for _, j := range arr {
			if max < j {
				max = j
			}
			if min > j {
				min = j
			}
		}

	}
	fmt.Print("output ")
	return (sum - min - max) / (len(arr) - 2)
}
