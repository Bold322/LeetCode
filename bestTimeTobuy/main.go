package main

import "fmt"

// type prod struct {
// 	start int
// 	end   int
// 	prof  int
// }

// func maxProfit(prices []int) int {
// 	profit := 0

// 	for i := 0; i < len(prices); i++ {
// 		profits := []prod{}
// 		realProfs := []prod{}

// 		start := i

// 		m := 0
// 		s := 0
// 		e := 0

// 		for k := i; k < len(prices); k++ {
// 			if prices[k] > prices[start] {
// 				profits = append(profits, prod{
// 					start: start,
// 					end:   k,
// 					prof:  prices[k] - prices[start],
// 				})
// 			}

// 			if len(profits) > 0 {
// 				for l := 0; l < len(profits); l++ {
// 					if m < profits[l].prof {
// 						m = profits[l].prof
// 						s = profits[l].start
// 						e = profits[l].end
// 					}
// 				}
// 			}

// 		}
// 		if len(profits) > 0 {
// 			realProfs = append(realProfs, prod{
// 				start: s,
// 				end:   e,
// 				prof:  m,
// 			})
// 		}

// 		profits = []prod{}

// 		if len(realProfs) > 0 {
// 			start = realProfs[0].end + 1

// 			m = 0
// 			s = 0
// 			e = 0

// 			for k := start; k < len(prices); k++ {
// 				if prices[k] > prices[start] {
// 					profits = append(profits, prod{
// 						start: start,
// 						end:   k,
// 						prof:  prices[k] - prices[start],
// 					})
// 				}

// 				if len(profits) > 0 {
// 					for l := 0; l < len(profits); l++ {
// 						if m < profits[l].prof {
// 							m = profits[l].prof
// 							s = profits[l].start
// 							e = profits[l].end
// 						}
// 					}
// 				}

// 			}
// 		}

// 		if len(profits) > 0 {
// 			realProfs = append(realProfs, prod{
// 				start: s,
// 				end:   e,
// 				prof:  m,
// 			})
// 		}

// 		for l := 0; l < len(realProfs); l++ {
// 			if profit < realProfs[l].prof {
// 				profit = profit + realProfs[l].prof
// 				s = realProfs[l].start
// 				e = realProfs[l].end
// 			}
// 		}

// 	}

// 	return profit
// }

// type prod struct {
// 	start int
// 	end   int
// 	prof  int
// }

// func maxProfit(prices []int) int {
// 	profit := 0

// 	for i := 0; i < len(prices); i++ {
// 		profits := []prod{}
// 		realProfs := []prod{}

// 		start := i

// 		m := 0
// 		s := 0
// 		e := 0

// 		for k := i; k < len(prices); k++ {
// 			if prices[k] > prices[start] {
// 				profits = append(profits, prod{
// 					start: start,
// 					end:   k,
// 					prof:  prices[k] - prices[start],
// 				})
// 			}

// 			if len(profits) > 0 {
// 				for l := 0; l < len(profits); l++ {
// 					if m < profits[l].prof {
// 						m = profits[l].prof
// 						s = profits[l].start
// 						e = profits[l].end
// 					}
// 				}
// 			}

// 		}
// 		if len(profits) > 0 {
// 			realProfs = append(realProfs, prod{
// 				start: s,
// 				end:   e,
// 				prof:  m,
// 			})
// 		}

// 		profits = []prod{}

// 		if len(realProfs) > 0 {
// 			start = realProfs[0].end + 1

// 			m = 0
// 			s = 0
// 			e = 0

// 			for k := start; k < len(prices); k++ {
// 				if prices[k] > prices[start] {
// 					profits = append(profits, prod{
// 						start: start,
// 						end:   k,
// 						prof:  prices[k] - prices[start],
// 					})
// 				}

// 				if len(profits) > 0 {
// 					for l := 0; l < len(profits); l++ {
// 						if m < profits[l].prof {
// 							m = profits[l].prof
// 							s = profits[l].start
// 							e = profits[l].end
// 						}
// 					}
// 				}

// 			}
// 		}

// 		if len(profits) > 0 {
// 			realProfs = append(realProfs, prod{
// 				start: s,
// 				end:   e,
// 				prof:  m,
// 			})
// 		}

// 		for l := 0; l < len(realProfs); l++ {
// 			if profit < realProfs[l].prof {
// 				profit = profit + realProfs[l].prof
// 				s = realProfs[l].start
// 				e = realProfs[l].end
// 			}
// 		}

// 	}

// 	return profit
// }

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit1 := make([]int, len(prices))
	profit2 := make([]int, len(prices))

	// for i := 0; i < len(prices); i++ {
	// 	profit1 = append(profit1, 0)
	// 	profit2 = append(profit2, 0)
	// }

	start := prices[0]
	for i := 1; i < len(prices); i++ {
		end := prices[i]
		fmt.Println("Start:", start, "--End:", end)
		if profit1[i-1] < end-start {
			profit1[i] = end - start
		} else {
			profit1[i] = profit1[i-1]
		}
		fmt.Println("Profit1--", i, ":", profit1[i])

		if start > end {
			start = end
		}
	}

	fmt.Println("TRANS2")
	start2 := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		end := prices[i]
		fmt.Println("Start:", start, "--End:", end)

		if profit2[i+1] < start2-end {
			profit2[i] = start2 - end
		} else {
			profit2[i] = profit2[i+1]
		}

		fmt.Println("Profit2--", i, ":", profit2[i])
		if start2 < end {
			start2 = end
		}
	}

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if maxProfit < profit1[i]+profit2[i] {
			maxProfit = profit1[i] + profit2[i]
		}
	}

	return maxProfit
}

func main() {
	prices := []int{2, 1, 4}

	fmt.Println(maxProfit(prices))
}
