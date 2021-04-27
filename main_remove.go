package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/awnzl/gnl/internal/gnl"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r1 := strings.NewReader("01234567890123456first1 string\nsecond1 string\nthird1 string\n")
	// r2 := strings.NewReader("first2 string\nsecond2 string\nthird2 string\n")

	fmt.Println(string(<-gnl.GetNextLine(ctx, r1)))
	// fmt.Println(string(<-gnl.GetNextLine(ctx, r2)))
	fmt.Println(string(<-gnl.GetNextLine(ctx, r1)))
	// fmt.Println(string(<-gnl.GetNextLine(ctx, r2)))
	fmt.Println(string(<-gnl.GetNextLine(ctx, r1)))
	// fmt.Println(string(<-gnl.GetNextLine(ctx, r2)))
	// fmt.Println(string(<-gnl.GetNextLine(ctx, r2)))
}
