package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
}

func id(x interface{}) interface{} {
	return x
}

func composition(a func(int) float32, b func(float32) string) func(int) string {
	return func(number int) string {
		return b(a(number))
	}
}
