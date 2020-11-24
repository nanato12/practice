# 関数

## 基本
他の言語と何ら変わりない
```go
x := f(10, 1+1, y)
```

## 組み込み関数

| 関数 | 説明 |
| :--- | :--- |
| print / println | 表示を行う |
| make | コンポジット型の初期化 |
| new | コンポジット型の初期化 |
| len / cap | 長さ / 容量を取得 |
| delete | マップから指定したキーのエントリを削除 |
| complex | 複素数型を作成 |
| imag / real | 複素数の虚部、実部を取得 |
| panic / recover | パニックを起こす / 回復 |

## 関数の定義
通常
```go
func add(x int, y int) int {
    return x + y
}
```
複数の戻り値  
引数が同じ型の場合、まとめて記述できる。
```go
func swap(x, y int) (int, int) {
    return y, x
}
```
名前付き戻り値  
戻り値に変数名を明示していれば`return`のみでよい。  
**通常は使われない**
```go
func swap(x, y int) (x2, y2 int) {
    y2, x2 = x, y
    return
}
```

## 多値の戻り値
```go
x, y := swap(10, 20)
```
使わない場合は _ (ブランク変数) を用いる
```go
x, _ := swap(10, 20)
_, y := swap(10, 20)
```

## 値の入れ替え
一時変数なしで値を入れ替えることができる。
```go
x, y = y, x
```

## 無名関数
クロージャ
```go
msg := "Hello, world!"
func() {
    println(msg)
}()
```

関数型
```go
fs := make([]func() string, 2) // 返戻値stringの関数型の容量2のスライス
// スライスの各要素にクロージャを詰める
fs[0] = func() string { return "hoge" }
fs[1] = func() string { return "fuga" }
// 各要素の関数を実行
for _, f := range fs {
    fmt.Println(f())
}
```
クロージャとよくあるバグ  
変数`i`を表示するクロージャを詰めて、実行する。  
2回目のforに入る時に`変数iは「2」になっている`ので実行結果は`2, 2, 2`
```go
fs2 := make([]func(), 3) // 返戻値なしの関数型の容量3のスライス
// 各要素に関数を詰める
for i := range fs2 {
    fs2[i] = func() { fmt.Println(i) }
}
// 各要素の関数を実行
for _, f := range fs2 {
    f()
}
```

## ポインタ
普通のポインタの概念と一緒
```go
func f(xp *int) {
	*xp = 100
}

var x2 int
fmt.Println(x2)
f(&x2)
fmt.Println(x2)
```

2階層以上のポインタ
```go
type T1 *int
type T2 *int
var np *T2
fmt.Println(np)
```

**内部でポインタを使っているデータ型**
コンポジット型の一部
- スライス
- マップ
- チャネル

```go
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
```
