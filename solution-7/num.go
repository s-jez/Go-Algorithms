package main

import "fmt"

func main() {
	result := NumInList([]int{1, 2, 3, 4, 5}, 5)
	fmt.Println(result)
}

func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
