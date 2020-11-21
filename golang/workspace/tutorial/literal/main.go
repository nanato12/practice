package main

import (
	"fmt"
	"strings"
)

func main() {
	// float64をintにキャスト
	var f float64 = 10
	var n int = int(f)
	fmt.Println(n)

	var ns []int
	var m map[string]int
	var person struct {
		name string
		age  int
	}
	fmt.Println(ns)
	fmt.Println(m)
	fmt.Println(person)

	satou := struct {
		name string
		age  int
	}{
		name: "satou",
		age:  10,
	}
	fmt.Println(satou)
	fmt.Println(satou.name)
	fmt.Println(satou.age)

	ns3 := [...]int{5: 1, 10: 5}
	fmt.Println(ns3)

	fmt.Println(ns3[3])
	fmt.Println(len(ns3))
	fmt.Println(ns3[5:7])

	// ゼロ値はnil
	var ns1 []int

	// 長さと容量を指定して初期化
	// 各要素はゼロ値で初期化される
	ns1 = make([]int, 3, 10)
	fmt.Println(ns1)

	// スライスリテラルで初期化
	// 要素数は指定しなくていい
	// 自動で配列は作られる
	var ns2 = []int{10, 20, 30, 40, 50}
	fmt.Println(ns2)

	// 5番目が50, 10番目が100のスライスを初期化
	ns4 := []int{5: 50, 10: 100}
	fmt.Println(ns4)

	ns5 := []int{10, 20, 30, 40, 50}
	fmt.Println(ns5[3])   // 40
	fmt.Println(len(ns5)) // 5
	fmt.Println(cap(ns5)) // 5
	// 容量が足りない場合は配列が再確保される
	ns5 = append(ns5, 60, 70)
	fmt.Println(ns5)      // [10 20 30 40 50 60 70]
	fmt.Println(len(ns5)) // 7
	fmt.Println(cap(ns5)) // 10

	a := []int{10, 20}
	b := append(a, 30)
	a[1] = 100
	c := append(b, 40)
	b[1] = 200
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	fmt.Println(cap(c))
	sliceSample()
}

// sliceSample ...
func sliceSample() {
	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4
	fmt.Println(ns[n:]) // [30 40 50]
	fmt.Println(ns[:m]) // [10 20 30 40]

	ms := ns[:m:m]
	fmt.Println(cap(ms)) //4
	mapSample()
}

// mapSample ...
func mapSample() {
	// ゼロ値はnil
	var m map[string]int
	//makeで初期化
	m = make(map[string]int)
	// 容量を指定できる
	m = make(map[string]int, 10)
	fmt.Println(m)

	// リテラルで初期化
	ma := map[string]int{"x": 100, "y": 200}
	fmt.Println(ma["x"])
	ma["z"] = 300
	fmt.Println(ma)

	// 存在確認
	n, ok := ma["z"]   // 存在する場合 value, bool
	fmt.Println(n, ok) // 300 true

	// キーを指定して削除
	delete(ma, "z")

	// 存在確認
	n, ok = ma["z"]    // 存在しない場合　ゼロ値, bool
	fmt.Println(n, ok) // 0 false

	// からの場合
	ma = map[string]int{}

	counts := map[string]int{}

	str := "dog dog dog cat dog"
	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	fmt.Println(counts)
}
