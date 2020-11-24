# パッケージ

## import
サードパーティのインポート  
まあ普通
```go
import (
    .....
)
```
変数文字
```go
import (
	"github.com/tenntenn/greeting"
    greetingv2 "github.com/tenntenn/greeting/v2"
)
```

## パッケージ外へのエクスポート
- 先頭大文字にした識別子がエクスポートされる
- 他のパッケージから利用できるようにする
```go
var Hoge string // エクスポートできる
var huga string // エクスポートできない
```

## ライブラリ
- main関数のないGoのプログラム
- エクスポートされたものを利用する
