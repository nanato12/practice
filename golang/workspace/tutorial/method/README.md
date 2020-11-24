# メソッド

レシーバと紐付けられた関数
- データとそれに対する操作を紐づけるために用いる
- ドットでメソッドにアクセスする

```go
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

var hex Hex = 100
fmt.Println(hex.String())
```

メソッドも値として扱える。
```go
f := hex.String
fmt.Println(f())
```

## レシーバ
メソッドに紐付けられた変数
- メソッド呼び出し時には通常の引数と同じような扱いになる
  - コピーが発生する。
- ポインタを用いることでレシーバへの変更を呼び出し元に伝えることができる
  - レシーバがポインタの場合もドットでアクセスする

```go
type T int

func (t *T) f() { fmt.Println("hi") }

var v T
(&v).f()
v.f()
```

ポインタ型のメソッドリスト  
*TはTのメソッドとして扱われる。
```go
package main

type T struct{}

func (t T) f()  {}
func (t *T) g() {}

func main() {
	(T{}).f()
	(&T{}).f()
	(*&T{}).f()

	// (T{}).g() // <- できない
	(&T{}).g()
	(*&T{}).g()
}
```

レシーバにできる型
- typeで定義した型
  - ユーザー定義型など
- ポインタ型
  - 内部に変更を加えたい時
- 内部にポインタを持つ型
  - マップ、スライスなど
