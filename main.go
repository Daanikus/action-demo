package main

import (
	"context"
	"fmt"
)

func main() {
	Thing()
	foo("Hello", context.Background())
}

func Thing() {
	var t []string
	fmt.Println(t)
}

func foo(msg string, ctx context.Context) {
	fmt.Println(msg)
}
