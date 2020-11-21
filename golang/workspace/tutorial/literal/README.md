# Composite, Literal

## キャスト
float64をintにキャスト
```go
var f float64 = 10
var n int = int(f)
```

## コンポジット型

スライスやマップはmakeなどで初期化が必要なためnil

| 種類 | 説明 | ゼロ値 |
| :--- | :--- | :--- |
| 構造体 | 型の異なるデータ型を集めたデータ型 | フィールドが全てゼロ値 |
| 配列 | 同じ型のデータを集めて並べたデータ型 | 要素が全てゼロ値 |
| スライス | 配列の一部を切り出したデータ型 | nil |
| マップ | キーと値をマッピングさせたデータ型 | nil |

## 型リテラル
- 型の具体的な定義を書き下した型の表現方法
- コンポジット型などを表現するために使う
- 変数定義、ユーザー定義などで使用する

```go
// int型のスライスの型リテラルを使った定義
var ns []int

// mapの型リテラルを使った定義
// keyがstring
// valueがint
var m map[string]int
```

## 構造体

var 変数名 型 (または型リテラル)
```go
var person struct {
    name string
    age  int
}
```
値を代入する。
```go
satou := struct {
    name string
    age  int
}{
    name: "satou",
    age:  10,
}
```
フィールドにアクセスする。
```go
fmt.Println(satou.name)
fmt.Println(satou.age)
```

## 配列

goの配列は`固定長配列`

- 要素の型は全て同じ
- 要素数が違えば別の型
- 要素数を変更できない
- 型は型リテラルで記述することが多い

要素数5の配列に要素数4を代入などはできない
```go
var ns [5]int
```

## 配列の初期化
```go
// ゼロ値で初期化
var ns [5]int
// 配列リテラルで初期化
var ns2 [5]int{1, 2, 3 4, 5}
// 要素数を推論
ns3 := [...]int{1, 2, 3, 4, 5}
// 5番目が1, 10番目が5, 他が0の要素数11の配列
ns3 := [...]int{5: 1, 10: 5} // [0 0 0 0 0 1 0 0 0 0 5]
```

コンポジットリテラル
https://golang.org/ref/spec#Composite_literals

## 配列の操作
```go
ns3 := [...]int{5: 1, 10: 5}

// 特定の要素にアクセス
fmt.Println(ns3[3])
// 配列の長さを取得
fmt.Println(len(ns3))
// 配列からスライスを取得
fmt.Println(ns3[5:7])
```

## スライス
配列の一部を取り出したデータ型
- 要素の方は全て同じ
- 要素数は型情報を含まない

```go
ns3 := [...]int{10, 20, 30, 40, 50}
ns[1:3] // [20 30]

len(ns[1:3]) // 2 (スライスの長さ)
cap(ns[1:3]) // 4 (拡張可能な数) 最初のポインタから最後までの数
```

## スライスの初期化
```go
// ゼロ値はnil
var ns1 []int

// 長さと容量を指定して初期化
// 各要素はゼロ値で初期化される
ns1 = make([]int, 3, 10)

// スライスリテラルで初期化
// 要素数は指定しなくていい
// 自動で配列は作られる
var ns2 = []int{10, 20, 30, 40, 50}

// 5番目が50, 10番目が100のスライスを初期化
ns4 := []int{5: 50, 10: 100}
```

## スライスと配列の関係
スライスはベースとなる配列が存在している

```go
ns := make([]int, 3, 10)
```
これは
```go
var array [10]int
ns := array[:3]
```
とほぼ同じ処理

```go
ms := []int{10, 20, 30, 40, 50}
```
これは
```go
// 配列リテラルで配列を作る
var array2 = [...]int{10, 20, 30, 40, 50}
// 0から5番目までのスライスを作る
ms := array[:5] //または array[:]
```

## スライスの操作
```go
// スライスリテラルで定義
ns5 := []int{10, 20, 30, 40, 50}

// 特定の要素にアクセス
fmt.Println(ns5[3]) // 40
// 長さ
fmt.Println(len(ns5)) // 5
// 容量
fmt.Println(cap(ns5)) // 5

// 要素の追加
// 容量が足りない場合は配列が再確保される
ns5 = append(ns5, 60, 70)
fmt.Println(ns5) // [10 20 30 40 50 60 70]

// 長さ
fmt.Println(len(ns5)) // 7
// 容量
fmt.Println(cap(ns5)) // 10
```

### appendの挙動
- 要素が足りる場合
  - 新しい要素をコピーする
  - lenを更新

- 要素が足りない場合
  - 元の2倍の容量の配列を確保し直す
  - 配列へのポインタを貼り直す
  - lenの更新
  - capの更新

```go
a := []int{10, 20} // 10番地の配列が生成される
fmt.Println(a, cap(a)) // [10 20] 2

b := append(a, 30) // 容量が足りないので容量が倍になった配列が20番地に生成される
a[1] = 100 // 10番地の配列の要素1に代入
fmt.Println(b, cap(b)) // [10 20 30] 4

c := append(b, 40) // 容量が足りるので20番地の配列のスライスが定義される
b[1] = 200 // 20番地の配列の要素1に代入
fmt.Println(c, cap(c)) // [10 200 30 40] 4
```

appendするスライスと受け取るスライスを別にすると予期しない挙動になる可能性があるので注意

## スライスへのスライス演算
```go
ns := []int{10, 20, 30, 40, 50}
n, m := 2, 4
fmt.Println(ns[n:]) // [30 40 50]
fmt.Println(ns[:m]) // [10 20 30 40]

ms := ns[:m:m]
fmt.Println(cap(ms)) //4
```

要素が100000ある配列の先頭10のスライスを使いまわすと、巨大な配列をコピーすることになる。  
その場合には, `:m:m`のようにトリムする

現在は`append`しかないので、go wikiみる。  
https://github.com/golang/go/wiki/SliceTricks

**スライスはnilとしか比較できない**
**配列は比較できる**

## マップ
- キーは **==で比較できる値**に限る
キーに出来ない値は...スライス  

構造体のvalueにスライスがあるとダメ  
基本的にスライスはポインタ操作なので。

```go
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
```

makeで初期化できるのは、**スライス、マップ、チャネル**  
マップの裏にも配列がいる。

**ワードカウント**  
ゼロ値があるので、valueは0で初期化
```go
counts := map[string]int{}

str := "dog dog dog cat dog"
for _, s := range strings.Split(str, " ") {
    counts[s]++
}
fmt.Println(counts) // map[cat:1 dog:4]
```

## ユーザー定義型

type 型名 基底型

```go
// 組み込み型を基にする
type MyInt int

// 他のパッケージの型を基にする
type MyWriter io.writer

// 型リテラルを基にする
type Person struct{
    Name string
}
```

基底型とユーザー定義型の相互キャストが可能
```go
type MyInt int
var n int = 100
m := MyInt(n)
n = int(m)
```

型なし定数から明示的なキャストは不要
```go
d := 10 * time.Second // time.Second は time.Duration型
// time.Duration型の基底型はint64
```
`基底型がint64であるtime.Duration型`と`定数10`は計算可能なのでキャスト不要

## 型エイリアス (Go 1.9以上)
- 完全に同じ型
- キャスト不要
- エイリアスの方ではメソッドの定義はできない
```go
type Applicant = http.Client
```

型名を出力する%Tが同じ元の型名を出す
```go
type Applicant = http.Client
func main () {
    fmt.Printf("%T", Applicant{}) // http.Client
}
