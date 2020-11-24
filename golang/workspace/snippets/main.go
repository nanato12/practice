package main

import "fmt"

const iota = 0

const (
	a = iota
	b
	c
)

func main() {
	false := true
	if false {
		fmt.Println(false)
	}

	type int string
	var n int
	fmt.Println(n + "sakura")

	fmt.Println(a, b, c)
}
