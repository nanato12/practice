package main

import "fmt"

func add(x int, y int) int {
	return y + x
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y int) (x2, y2 int) {
	x2, y2 = y, x
	return
}

func f(xp *int) {
	*xp = 100
}

func main() {
	x, y := 10, 20
	fmt.Println(x, y)

	fmt.Println(add(x, y))
	fmt.Println(swap(x, y))
	fmt.Println(swap2(x, y))

	msg := "Hello, world!"
	func() {
		println(msg)
	}()

	fs := make([]func() string, 2) // 返戻値stringの関数型の容量2のスライス
	// スライスの各要素にクロージャを詰める
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }
	// 各要素の関数を実行
	for _, f := range fs {
		fmt.Println(f())
	}

	fs2 := make([]func(), 3) // 返戻値なしの関数型の容量3のスライス
	// 各要素に関数を詰める
	for i := range fs2 {
		fs2[i] = func() { fmt.Println(i) }
	}
	// 各要素の関数を実行
	for _, f := range fs2 {
		f()
	}

	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}
	p2 := p
	p2.age = 20
	fmt.Println(p.age, p.name)
	fmt.Println(p2.age, p2.name)

	var x2 int
	fmt.Println(x2)
	f(&x2)
	fmt.Println(x2)

	type T1 *int
	type T2 *int
	var np *T2
	fmt.Println(np)

	var ns [][][][][][][][][][]int
	fmt.Println(ns)

	ns2 := []int{10, 20, 30}
	ns3 := ns2
	ns2[1] = 200
	fmt.Println(ns2[0], ns2[1], ns2[2]) // 10 200 30
	fmt.Println(ns3[0], ns3[1], ns3[2]) // 10 200 30

	ns4 := [...]int{10, 20, 30}
	ns5 := ns4
	ns4[1] = 200
	fmt.Println(ns4[0], ns4[1], ns4[2]) // 10 200 30
	fmt.Println(ns5[0], ns5[1], ns5[2]) // 10 20 30
}
