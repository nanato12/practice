# variable

変数のいろいろ

nullは`nil`

## 型の定義
```go
// 型指定と代入
var n int = 100

// 型指定してから代入
var n int
n = 100

// 型指定を省略
var n = 100

// varと型指定を省略
n := 100

// 複数
var (
    n = 100
    m = 200
)
```

## 組み込み型

| 意味 | 型 |
| :--- | :--- |
| 整数 | int, int8, int16, int32, int64 <br> uint, unit8, uint16, uint32, uint64 <br> uintptr, byte, rune |
| 浮動小数点数 | float32, float64 |
| 複素数 | complex64, complex128 |
| 文字列 | string |
| 真偽値 | bool |
| エラー | error |

uintptr: ポインタの加算、減算するときに使う  
byte: int8と同義  
rune: unicodeの1コードポイント

## 変数のゼロ値

| 型 | ゼロ値 |
| :--- | :--- |
| int, float, etc... | 0 |
| string | "" |
| bool | false" |
| error | nil |
