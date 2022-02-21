package main

import "fmt"

func main() {
	number1, number2, number3 := numbers()

	fmt.Println(number1, number2, number3)
}
func numbers() (int, int, int) {
	return 1000, 10, 100
}
