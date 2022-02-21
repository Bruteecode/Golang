package main

import "fmt"

func main() {
	string1, string2, integer1, integer2 := numbers()

	fmt.Println(string1, string2, integer1, integer2)

}
func numbers() (string, string, int, int) {
	return "akash", "chopra", 7, 212
}
