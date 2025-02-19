package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// this do not work, i am trying to pass a variable when i have to pass an adress(pointer)
	// zeroptr(i)
	// fmt.Println("no pass &i", i)

	fmt.Println("pointer:", &i)
}
