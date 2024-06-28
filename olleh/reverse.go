package olleh

import (
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %v, j: %v\n", i, j)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
