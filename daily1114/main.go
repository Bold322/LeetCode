package main

import (
	"fmt"
	"sort"
)

func minimizedMaximum(n int, quantities []int) int {

	sort.Slice(quantities, func(i, j int) bool {
		return quantities[i] < quantities[j]
	})

	if len(quantities) == n {
		return quantities[len(quantities)-1]
	}

	s := n % len(quantities)
	f := n / len(quantities)
	min := quantities[0]
	max := quantities[len(quantities)-1]
	start := min / f
	if s > 0 {
		return max - (min / (f - 1))
	}

	end := max - (start * (f + s))

	fmt.Println("asdfasdf")


	

	// start := quantities[len(quantities)-1]
	// if quantities[len(quantities)-1]%2 > 0 {
	// 	start = quantities[len(quantities)-1] + 1
	// }

	// fmt.Println("Start:", start)
	// fmt.Println("s:", s)
	// fmt.Println("f:", f)

	// if s > 0 {
	// 	return start - (quantities[0] * f)
	// }

	return end
}

func main() {
	n := 6
	m := []int{11, 6}

	fmt.Println(minimizedMaximum(n, m))
}
