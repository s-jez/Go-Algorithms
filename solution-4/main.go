package main

import "fmt"

func main() {
	FizzBuzz(100)
}
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		printFizzBazzValue(i)
		if i != n {
			fmt.Print(", ")
		}
	}
}
func printFizzBazzValue(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("Fizz Buzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Print(n)
	}
}
