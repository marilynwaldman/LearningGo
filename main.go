package main

import (
	"fmt"
	src "./src"
	"time"

	"strings"
)
func greet(name string) {
	msg := name + fmt.Sprintf(" you arrived at (at %v)", time.Now().String())

	closure := func() {
		fmt.Printf("Hey %s!n", msg)
	}
	closure()

}

func main() {
	//fib := src.MemorizedFib;
	//fmt.Printf("Fib: %w\n", fib(8))
	//greet("Alice")
	//src.FibSimple(5)
	//fmt.Printf("Fib: %w\n", src.FibSimple(8))
	var recognize = func(name string) string {
		return fmt.Sprintf("Hey %s", name)
	}
	var emphasize = func(statement string) string {
		return fmt.Sprintf(strings.ToUpper(statement) + "!")
	}


	var greetFoG = src.Compose(recognize, emphasize)
	fmt.Println(greetFoG("Gopher"))


}

