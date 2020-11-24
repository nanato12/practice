package main

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting"
	greetingv2 "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do())
	fmt.Println(greetingv2.Do(time.Now()))
	fmt.Println(time.Now())
}
