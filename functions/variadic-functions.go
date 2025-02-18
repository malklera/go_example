package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/*
This do not work, ... can only go with the last argument of the function
func rest(nums ...int, words ...string) {
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ", words[i], "\n")
	}
}
*/

func rest(nums []int, words []string) {
	if len(nums) != len(words) {
		fmt.Println("Error, nums and words has to be same length")
		return
	}

	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ", words[i], "\n")
	}
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	words := []string{"a", "b", "c", "d"}
	rest(nums, words)
}
