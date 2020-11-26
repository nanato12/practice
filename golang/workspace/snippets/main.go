package main

import (
	"fmt"
	"strconv"
)

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

	PostID := "051a72fc"
	if postId, err := strconv.ParseInt(PostID, 16, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(postId)
	}
}
