package main

import "fmt"

func shortestSubarray(nums []int, k int) int {
	answers := []int{}
	ans := len(nums)

	n := len(nums)

	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			for s := n - 1; s >= end; s-- {
				p := 0
				for i := s; i >= end; i-- {
					p = p + nums[i]
				}
				if p == k {
					answers = append(answers, s-end+1)
				}
			}
		}
	}

	if len(answers) < 1 {
		return -1
	}

	for i := range answers {
		ans = min(answers[i], ans)
	}

	return ans
}

func main() {

	nums := []int{48,99,37,4,-31}
	k := 140

	fmt.Println(shortestSubarray(nums, k))
}
