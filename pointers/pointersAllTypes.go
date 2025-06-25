package main

import "fmt"

func zeroVal(aVal bool, bVal int, cVal string, dVal [3]int, eVal []string,
	fVal map[string]int, gVal myStruct) {
	aVal = false
	bVal = 5
	cVal = "five"
	dVal = [...]int{4, 5, 6}
	// WARN: this changed the original value
	eVal[0] = "four"
	fVal["two"] = 4
	gVal.q[0] = "seven"
}

func zeroPtr(aVal *bool, bVal *int, cVal *string, dVal *[3]int, eVal *[]string,
	fVal *map[string]int, gVal *myStruct) {
	*aVal = false
	*bVal = 7
	*cVal = "six"
	*dVal = [...]int{7, 8, 9}
	(*eVal)[1] = "five"
	(*fVal)["three"] = 5
	(*gVal).q[2] = "four"
}

type myStruct struct {
	m bool
	n int
	o string
	p [3]int
	q []string
}

func main() {
	var a bool = true
	var b int = 10
	var c string = "ten"
	var d = [3]int{1, 2, 3}
	var e = []string{"one", "two", "three"}
	var f = map[string]int{"one": 1, "two": 2, "three": 3}
	var g = myStruct{
		m: true,
		n: 10,
		o: "ten",
		p: [3]int{1, 2, 3},
		q: []string{"one", "two", "three"},
	}
	fmt.Println("initial a:", a)
	fmt.Println("initial b:", b)
	fmt.Println("initial c:", c)
	fmt.Println("initial d:", d)
	fmt.Println("initial e:", e)
	fmt.Println("initial f:", f)
	fmt.Println("initial g:", g)

	fmt.Println("\ncalling zeroVal\n")
	zeroVal(a, b, c, d, e, f, g)

	fmt.Println("after zeroVal a:", a)
	fmt.Println("after zeroVal b:", b)
	fmt.Println("after zeroVal c:", c)
	fmt.Println("after zeroVal d:", d)
	fmt.Println("after zeroVal e:", e)
	fmt.Println("after zeroVal f:", f)
	fmt.Println("after zeroVal g:", g)

	fmt.Println("\ncalling zeroPtr\n")
	zeroPtr(&a, &b, &c, &d, &e, &f, &g)

	fmt.Println("after zeroPtr a:", a)
	fmt.Println("after zeroPtr b:", b)
	fmt.Println("after zeroPtr c:", c)
	fmt.Println("after zeroPtr d:", d)
	fmt.Println("after zeroPtr e:", e)
	fmt.Println("after zeroPtr f:", f)
	fmt.Println("after zeroPtr g:", g)
}
