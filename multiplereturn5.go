package main

import "fmt"

func main() {
	c := player("MS")

	fmt.Println(c.describe("is a legend"))
}

type player string

func (c player) describe(description string) string {
	return string(c) + " " + description
}
