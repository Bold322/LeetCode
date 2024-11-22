package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 3, 2, 5}
	k := 3

	fmt.Println(resultsArray(n, k))

}

func resultsArray(nums []int, k int) []int {

	lg := len(nums)
	rs := make([]int, lg-k+1)

	if lg == 1 || k == 1 {
		return nums
	}

	for i := 0; i < lg-k+1; i++ {
		for j := i + 1; j < i+k; j++ {
			if nums[j]-1 == nums[j-1] {
				rs[i] = max(rs[i], nums[j])
			} else {
				rs[i] = -1
				break
			}
		}
	}

	return rs
}
