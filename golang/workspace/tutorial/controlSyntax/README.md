# Control syntax
制御構文のあれこれ

## 条件分岐 if
条件式に()はかかない
```go
x := 100
if x == 1 {
    fmt.Println("xは1")
} else if x == 2 {
    fmt.Println("xは2")
} else {
    fmt.Println(x)
}
```

{}の改行はできない。  
{}の省略はできない。
```go
// コンパイルエラー
if x == 0
{
}

// コンパイルエラー
if (x == 0) fmt.Println(x)
```

## 条件分岐 switch
**breakがいらない**ってまじか  
次のcase、defaultを評価したい時は`fallthrough`

```go
switch x {
case 1:
    fmt.Println("xは1")
case 2:
    fmt.Println("xは2")
case 100, 200:
    fmt.Println("xは100か200")
    fallthrough // またぐ
default:
    fmt.Println(x)
}
```

**caseに式が使える**ってまじか  
ifの代わりに使える。  
めっちゃ読みやすい。
```go
switch {
case x == 1:
    fmt.Println("xは1")
case x == 2:
    fmt.Println("xは2")
case x == 100:
    fmt.Println("xは100")
}
```

## 繰り返し for
普通の繰り返し
```go
// 初期値; 継続条件; 更新分
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
whileのように
```go
// 継続条件のみ
i := 0
for i := 0; i < 10; i++ {
    fmt.Println(i)
    i++
}
```
無限ループ while true
```go
for {
    fmt.Println("hello")
}
```
rangeを使った繰り返し  
第一返戻値がindex (0, 1, 2)
第二返戻値がvalue (1, 2, 3)
```go
for _, value := range []int{1, 2, 3} {
    fmt.Println(value)
}
```

繰り返しの抜け出し  
`break`使う
```go
var i int = 1
for {
    if i%2 == 0 {
        fmt.Println(i)
        break
    }
    i++
}
```
ラベル指定の`break`  
switch内にbreakがあるとswitchの`break`と混合するため、ラベルを指定して処理する。  
`break LABEL`
```go
LOOP:
	for {
		switch {
		case i >= 10:
			fmt.Println(i)
			break LOOP
		}
		i++
	}
```
