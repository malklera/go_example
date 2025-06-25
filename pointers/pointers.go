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

	fmt.Print("\n\n")

	name := "leopoldo" // declare a variable
	showValue(&name)   // pass a pointer to a variable
	// the following give error
	// showValue(name)
	// showValue(*name)

	testSlice := []string{"a", "b", "c", "d"}
	fmt.Println("testSlice[0]:", testSlice[0])
	showSlice(&testSlice)
	fmt.Println("testSlice[0]:", testSlice[0])
}

func showSlice(testSlice *[]string) {
	// func showSlice(testSlice []*string) {
	// func showSlice(testSlice *([]string)) {
	// func showSlice(testSlice *string) {
	fmt.Println("showSlice")
	// fmt.Println("testSlice[0]:", testSlice[0])
	fmt.Println("testSlice:", testSlice)
	fmt.Println("*testSlice:", *testSlice)
	fmt.Println("(*testSlice)[0]:", (*testSlice)[0])
	// fmt.Println("&testSlice[0]:", &testSlice[0])
	(*testSlice)[0] = "e"
}

func showValue(name *string) { // recibe a pointer
	fmt.Println("showValue")
	fmt.Println("name:", name)   // pointer address
	fmt.Println("*name:", *name) // value on that adress

	printValue(name, "pepe") // pass a pointer
	// name = "teo"		this give error because i try to asign string to adress
	*name = "jhon"
	printValue(name, "jakob")
	// the following give error
	// printValue(*name, "mujica")
	// printValue(&name, "fedor")
}

func printValue(name *string, newName string) {
	fmt.Println("printValue")
	fmt.Println("name:", name)
	fmt.Println("*name:", *name)
	*name = newName
	// the folowing give error
	// name = newName
	// &name = newName

	fmt.Println("name:", name)
	fmt.Println("*name:", *name)
	fmt.Println("\n")
}
