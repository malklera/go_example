package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points, this would not print
	// "0" "g" "\n" "1" "o" ... it will print "0" "unicode rune"(g=103)
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for i, c := range [6]string{"g", "o", "l", "a", "n", "g"} {
		fmt.Println(i, c)
	}
}
