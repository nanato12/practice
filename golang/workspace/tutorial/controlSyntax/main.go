package main

import "fmt"

func main() {
	x := 100
	if x == 1 {
		fmt.Println("xは1")
	} else if x == 2 {
		fmt.Println("xは2")
	} else {
		fmt.Print("xは")
		fmt.Println(x)
	}

	switch x {
	case 1:
		fmt.Println("xは1")
	case 2:
		fmt.Println("xは2")
	case 100, 200:
		fmt.Println("xは100か200")
		fallthrough // またぐ
	default:
		fmt.Println(x)
	}

	switch {
	case x == 1:
		fmt.Println("xは1")
	case x == 2:
		fmt.Println("xは2")
	case x == 100:
		fmt.Println("xは100")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int = 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for _, value := range []int{1, 2, 3} {
		fmt.Println(value)
	}

	i = 1
	for {
		if i >= 10 {
			fmt.Println(i)
			break
		}
		i++
	}

LOOP:
	for {
		switch {
		case i >= 10:
			fmt.Println(i)
			break LOOP
		}
		i++
	}
}
