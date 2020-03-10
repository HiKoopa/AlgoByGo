package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	a := num2
	for a > 0 {
		fmt.Println(num1 * (a % 10))
		a /= 10
	}
	fmt.Println(num1 * num2)
}