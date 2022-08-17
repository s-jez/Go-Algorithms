package main

import "fmt"

func main() {
	result := Sum([]int{3, 3, 3})
	fmt.Println(result)
}
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])

}

Sum({2, 3, 4}) = 2 + Sum({}) // 0 
