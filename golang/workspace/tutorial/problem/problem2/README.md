# problem2

次のプログラムで`true`になるパターン

```go
package main

func main() {
	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
```
真偽値表を作れってさ

| a | b | c | a && b | !c | a && b \|\| !c |
| :--- | :--- | :--- | :--- | :--- | :--- |
| false | false | false | false | true | true
| false | false | true | false | false | false
| false | true | false | false | true | true
| false | true | true | false | false | false
| true | false | false | false | true | true
| true | false | true | false | false | false
| true | true | false | true | true | true
| true | true | true | true | false | true

