package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
