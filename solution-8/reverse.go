package main

import (
	"fmt"
)

func main() {
	res := Reverse("Hello World!")
	fmt.Println(res)
}
func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}
