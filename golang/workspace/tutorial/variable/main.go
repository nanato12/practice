package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)

	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	n = 0x64
	fmt.Println(n)

	n = 0b0010
	fmt.Println(n)

	n = 5_000_000_000
	fmt.Println(n)
}
