package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// for anonymous functions i have to use var
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

	/*
		This do not work, when declaring a anonymous function i have to define the
		types it take and what it returns(if any)

		var fob func()

		fob = func(m string) string {
			return m + "pie"
		}
		fmt.Println(fob("apple"))
	*/
}
