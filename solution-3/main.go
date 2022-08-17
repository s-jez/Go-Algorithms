package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Reverse("cat"))      // => tac
	fmt.Println(ReverseSB("tomato")) // => otamot
}
func Reverse(word string) string {
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		res = res + string(word[i])
	}
	return res
}

// with StringBuilder
func ReverseSB(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
func ReverseMethod(word string) string {
	var res string
	for i := 0; i < len(word); i++ {
		res = string(word[i]) + res
	}
	return res
}
