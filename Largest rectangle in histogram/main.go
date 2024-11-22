package main

import (
	"fmt"
)

type Rectangle struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	// stack := []Rectangle{}
	maxValue := 0
	for start := 0; start < len(heights); start++ {
		minHeight := heights[start]
		for end := start; end < len(heights); end++ {
			if heights[end] < minHeight {
				minHeight = heights[end]
			}

			maxValue = max((end-start+1)*minHeight, maxValue)

			if minHeight*(len(heights)-start) <= maxValue {
				break
			}
		}
	}

	return maxValue
}

func asdf(heights []int) int {
	maxArea := 0
	stack := []Rectangle{}

	for currentIndex, currentHeight := range heights {
		startIndex := currentIndex

		// Calculate area for any taller rectangles in the stack that must end here
		for len(stack) > 0 && stack[len(stack)-1].height > currentHeight {
			// Pop the last rectangle and calculate the area it can form
			lastRectangle := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := currentIndex - lastRectangle.index
			area := lastRectangle.height * width
			if area > maxArea {
				maxArea = area
			}

			// Update startIndex to the popped rectangle's start for future rectangles
			startIndex = lastRectangle.index
		}

		// Add the current rectangle to the stack
		stack = append(stack, Rectangle{index: startIndex, height: currentHeight})
	}

	// Process any remaining rectangles in the stack
	for _, rectangle := range stack {
		width := len(heights) - rectangle.index
		area := rectangle.height * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {

	list := []int{1, 9, 3, 4}

	fmt.Println(largestRectangleArea(list))
}
