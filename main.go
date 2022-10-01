package main

import (
	"fmt"
)

// function add
func add(x int, y int) int {
	return x + y
}

// function swap
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	name := "Go Developers"
	fmt.Println("Welcome to go ", name)
	fmt.Println(add(42, 13))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

}
