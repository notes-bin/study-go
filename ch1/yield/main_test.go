package yield

import (
	"fmt"
	"testing"
	"iter"
)

func Backward[E any](s []E) func(yield func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			fmt.Printf("yield -> %#v\n", yield(i, s[i]))
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func TestYield(t *testing.T) {
	testdata := []string{"a", "b", "c"}
	for i, v := range Backward(testdata) {
		t.Logf("%#v, %#v\n", i, v)
	}
}

func f2(yield func(int, string) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i, fmt.Sprintf("I'm %d ", i)) {
			return
		}
	}
}

// for k := range f1 {
//
// }
func f1(yield func(int) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i) {
			return
		}
	}
}

// for range functionWithReturnsIsZero {
//
// }
func f0(yield func() bool) {
	for i := 0; i < 10; i++ {
		if !yield() {
			return
		}
	}
}

func TestIter(t *testing.T) {
	// 1. Basic usage, accepts iterator functions:
	// func(func() bool)
	// func(func(K) bool)
	// func(func(K, V) bool)
	fmt.Println("Test basic usage: for k,v := range f")
	for k, v := range f2 {
		fmt.Println(k, v)
	}

	fmt.Println("Test basic usage: for k := range f ")
	for k := range f1 {
		fmt.Println(k)
	}

	fmt.Println("Test basic usage: for range f ")
	for range f0 {
	}
}