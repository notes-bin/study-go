package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	for i := range ticker.C {
		fmt.Println(i)
	}
}
