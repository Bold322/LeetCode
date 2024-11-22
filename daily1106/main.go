package main

import (
	"fmt"
	"slices"
)

func canSortArray(nums []int) bool {

	if slices.IsSorted(nums) {
		return true
	} else {
		return false
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(canSortArray(s))
}
