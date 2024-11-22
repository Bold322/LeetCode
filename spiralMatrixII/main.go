package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	num := 0
	row := 0

	for row <= n/2-n%2+1 {

		for col := row; col < n-row; col++ {
			num++
			res[row][col] = num

		}

		for col := row + 1; col < n-row; col++ {
			num++
			res[col][n-row-1] = num
		}

		for col := n - row - 2; col > row-1; col-- {
			num++

			res[n-row-1][col] = num
		}

		for col := n - row - 2; col > row; col-- {
			num++

			res[col][row] = num
		}
		row++
	}

	return res
}

func main() {

	n := 5

	for i := 0; i < n; i++ {
		fmt.Println(generateMatrix(n)[i])
	}

}
