package main

import "fmt"

func main() {
	// 型のある定数
	const n int = 100
	// 型のない定数
	const m = 100
	// 定数式の使用
	const s = "Hello, world"
	// まとめる
	const (
		x = 100
		y = 200
	)

	// 巨大な数字を扱う
	const a = 100000000000000000000 / 100000000000000000000
	// bはint型になる
	var b = a
	fmt.Println(b)
}
