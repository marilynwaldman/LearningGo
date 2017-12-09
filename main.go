package main

import (
	"fmt"
	src "./src"
)

func main() {
	fib := src.MemorizedFib;
	fmt.Printf("Fib: %w\n", fib(8))
}

