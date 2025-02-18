package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)
	// check if a key exist
	_, v4 := m["k3"]
	fmt.Println("k3:", v4)

	_, v5 := m["k2"]
	fmt.Println("k2:", v5)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	clear(m)
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
