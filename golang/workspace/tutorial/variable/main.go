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
}
