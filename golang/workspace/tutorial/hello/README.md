# Hello, World !

```go
package main

import fmt "fmt"

func main() {
    fmt.Printf("Hello, world!\n")
}

// コメント

/*
コメント
*/
```

**package**を使ってソースファイルがどのパッケージに属しているかを必ず宣言する。

別のパッケージに含まれているものを使用するときは**import**を使う。

関数は`func`  
mainパッケージ内のmain関数が、まず最初(初期化処理後)に実行される。

コメントは`//`, `/**/`
