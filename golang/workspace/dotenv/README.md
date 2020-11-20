# dot env

`.env`に環境変数を保持して値を取ってくる。

```sh
$ go get github.com/joho/godotenv
```

`godotenv.Load()`でカレントディレクトリの`.env`を読み込んでくれる。  
あとは`os.Getenv('ENV')`で環境変数取ってくるだけ。

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvLoad load .env
func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	EnvLoad()
	message := fmt.Sprintf("user_name=%s password=%s", os.Getenv("USER_NAME"), os.Getenv("PASSWORD"))
	fmt.Println(message)
}
```
