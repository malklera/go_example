package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Print("type: ", reflect.TypeOf(f), ": ")
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Print("type: ", reflect.TypeOf(i), ": ")
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Print("type: ", reflect.TypeOf(d), ": ")
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Print("type: ", reflect.TypeOf(u), ": ")
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Print("type: ", reflect.TypeOf(k), ": ")
	fmt.Println(k)

	_, e := strconv.Atoi("wait")
	fmt.Println(e)
}
