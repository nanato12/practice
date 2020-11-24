package main

import "fmt"

type MyInt int

func (n *MyInt) inc() { *n++ }

func main() {
	var n MyInt
	fmt.Println(n)
	n.inc()
	fmt.Println(n)
}
