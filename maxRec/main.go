package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxArea := 0
	heights := make([]int, cols)

	for top := 0; top < rows; top++ {
		for col := 0; col < cols; col++ {
			if matrix[top][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}

		for start := 0; start < len(heights); start++ {
			minHeight := heights[start]
			for end := start; end < len(heights); end++ {
				if heights[end] < minHeight {
					minHeight = heights[end]
				}
				area := minHeight * (end - start + 1)
				if area > maxArea {
					maxArea = area
				}

				if minHeight*(len(heights)-start) <= area {
					break
				}
			}
		}
	}

	return maxArea
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalRectangle(matrix))
}
