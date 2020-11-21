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
| <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span>
| <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span>
| <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span>
| <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span>
| <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span>
| <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span> | <span style="color: red; ">false</span>
| <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span>
| <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: cyan; ">true</span> | <span style="color: red; ">false</span> | <span style="color: cyan; ">true</span>

