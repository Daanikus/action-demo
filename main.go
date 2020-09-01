package main

import (
	"context"
	"fmt"
)

type Bar struct {
	ID    int    `json:"id"`
	Name  string `yaml:"name"`
	Thing string `json:"name yaml:"name"`
}

func main() {
	Thing()
	foo("Hello", context.Background())
}

func Thing() {
	t := []string{}
	fmt.Println(t)
}

func foo(msg string, ctx context.Context) {
	fmt.Println(msg)
}
