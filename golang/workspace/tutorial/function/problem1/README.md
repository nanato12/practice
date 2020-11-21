# problem1

通らない次のプログラムを直す

```go
package problem1

import "fmt"

func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		fmt.Println("good")
	}
}
```

sumがintでavgがint/intだから型推論でintになる。  
avgがintだから4.5と比較不可  
sumにfoat64のキャストをかけて回避
