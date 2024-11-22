package main

import "fmt"

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	maxPointsOnLine := 1

	for i := 0; i < len(points); i++ {
		slopeCount := make(map[[2]int]int)
		duplicatePoints := 0
		currentMax := 0

		for j := i + 1; j < len(points); j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicatePoints++
			} else {
				g := dx
				for k := dy; k != 0; {
					g, k = k, g%k
				}

				if g < 0 {
					g = -g
				}
				dx /= g
				dy /= g

				if dx < 0 {
					dx = -dx
					dy = -dy
				} else if dx == 0 && dy < 0 {
					dy = -dy
				}

				key := [2]int{dy, dx}
				count := slopeCount[key]
				slopeCount[key] = count + 1
				currentCount := slopeCount[key]
				if currentCount > currentMax {
					currentMax = currentCount
				}
			}

			if currentMax+duplicatePoints+1 > maxPointsOnLine {
				maxPointsOnLine = currentMax + duplicatePoints + 1
			}
		}
	}

	return maxPointsOnLine
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 6}, {-1, -1}}
	fmt.Println(maxPoints(points))
}
