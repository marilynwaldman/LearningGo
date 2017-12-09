package main

import (
	"fmt"
	src "./src"
	"time"

)
func greet(name string) {
	msg := name + fmt.Sprintf(" you arrived at (at %v)", time.Now().String())

	closure := func() {
		fmt.Printf("Hey %s!n", msg)
	}
	closure()

}

func main() {
	fib := src.MemorizedFib;
	fmt.Printf("Fib: %w\n", fib(8))
	greet("Alice")


}

