# Problem4

レシーバに変更を与える。  
次のプログラムの実行結果が`0 1`になるように変更を加える。
```go
package main

import "fmt"

type MyInt int

func (n MyInt) inc() { n++ }

func main() {
	var n MyInt
	fmt.Println(n)
	n.inc()
	fmt.Println(n)
}
```
