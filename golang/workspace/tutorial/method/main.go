package main

import (
	"fmt"
)

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type T int

func (t *T) f() { fmt.Println("hi") }

func main() {
	var hex Hex = 100
	fmt.Println(hex.String())

	var v T
	(&v).f()
	v.f()

	f := hex.String
	fmt.Println(f())
}
