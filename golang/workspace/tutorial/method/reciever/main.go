package main

type T struct{}

func (t T) f()  {}
func (t *T) g() {}

func main() {
	(T{}).f()
	(&T{}).f()
	(*&T{}).f()

	// (T{}).g() // <- できない
	(&T{}).g()
	(*&T{}).g()
}
