package main

import (
	"fmt"
	"sort"
)

func primeSubOperation(nums []int) bool {

	m := 0
	for i := 0; i < len(nums)-1; i++ {
		for l := nums[i] - 1; l >= 0; l-- {
			div := 0

			for j := l; j > 0; j-- {
				if l%j == 0 {
					div++
					if div > 2 {
						break
					}
				}
			}
			if div <= 2 && l != 1 {
				if nums[i]-l > m {
					m = nums[i] - l
					break
				}
			}
		}
		nums[i] = m

		if m >= nums[i+1] {
			return false
		}

		if sort.SliceIsSorted(nums, func(i, j int) bool { return nums[i] <= nums[j] }) {
			return true
		}

	}

	return true
}

func primeSubOperations(nums []int) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}

	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			l := nums[i] - (nums[i+1] - 1)
			h := nums[i]
			p := getPrime(primes, l, h)
			if p == h {
				return false
			} else {
				nums[i] -= p
			}
		}
	}
	return true
}

func getPrime(primes []int, l, h int) int {
	for _, prime := range primes {
		if l <= prime && prime < h {
			return prime
		}
	}
	return h
}

func main() {
	n := []int{3, 3, 5, 5}

	fmt.Println(primeSubOperations(n))
}
