package main

import (
	"fmt"
)

func candy(ratings []int) int {

	candies := make([]int, len(ratings))

	candy := 0

	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := range candies {
		candy = candy + candies[i]
	}

	return candy
}

func main() {

	n := []int{1, 3, 2, 2, 1}

	fmt.Println(candy(n))
}
