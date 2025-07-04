package main

import (
	"fmt"
	"slices"
)

// this is the syntax to specify that you receive a slice
func sum(nums []int) {
	fmt.Println(nums)
}

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)

	for i := range(3) {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	slice1 := []int{1, 2, 3, 4}

	sum(slice1)

	testSlice := []string{"a", "b", "c", "d"}

	fmt.Println("testSlice[0]:", testSlice[0])
	fmt.Println("testSlice[len(testSlice)-1]:", testSlice[len(testSlice)-1])
	// the following panic index out of range
	// fmt.Println("testSlice[len(testSlice)]:", testSlice[len(testSlice)])

	var testAppend []string
	fmt.Println("\n\ntestAppend:", testAppend)
	testAppend = append(testAppend, "a")
	fmt.Println("testAppend = append(testAppend, a):", testAppend)

	secondSlice := []string{"b", "c", "d"}
	fmt.Println("secondSlice:", secondSlice)
	testAppend = append(testAppend, secondSlice...)
	fmt.Println("testAppend = append(testAppend, secondSlice...)", testAppend)
}
