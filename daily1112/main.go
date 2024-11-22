package main

import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	answer := []int{}

	sort.Slice(items, func(i, j int) bool {
		return items[i][1] > items[j][1]
	})

	if len(queries) < 1 {
		return []int{0}
	}
	for k := 0; k < len(queries); k++ {
		max := 0
		for i := 0; i < len(items); i++ {
			if queries[k] >= items[i][0] && max <= items[i][1] {
				max = items[i][1]
				break
			}
		}
		answer = append(answer, max)

	}

	return answer
}

func main() {

	items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	queries := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(maximumBeauty(items, queries))

}
