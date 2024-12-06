package main

import (
	"fmt"
)

func main() {
	grow := Exp(2)
	for i := range 10 {
		fmt.Printf("2^%d=%d\n", i, grow())
	}
	for f := range Fibonacci(10) {
		fmt.Println(f)
	}
}

func Exp(n int) func() int {
	e := 1
	return func() int {
		temp := e
		e *= n
		fmt.Println("e->", e, "n->", n)
		return temp
	}
}

func Fibonacci(n int) func(yield func(int) bool) {
	a, b, c := 0, 1, 1
	return func(yield func(int) bool) {
		for range n {
			fmt.Println("n->", n, "a->", a)
			if !yield(a) {
				return
			}
			a, b = b, c
			c = a + b
		}
	}
}
