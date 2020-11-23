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

## 多値の戻り値
```go
x, y := swap(10, 20)
```
使わない場合は _ (ブランク変数) を用いる
```go
x, _ := swap(10, 20)
_, y := swap(10, 20)
```
