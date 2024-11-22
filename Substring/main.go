package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	high := 0
	runes := []rune(s)
	start := 0
	for i := 0; i < len(s); i++ {
		for j := start; j < len(s); j++ {
			if runes[start] != runes[j] {
				high++
			}
		}
	}

	return high
}

// func lengthOfLongestSubstring(s string) int {
//     n := len(s)
//     maxLength := 0
//     lastIndex := make([]int, 128)

//     start := 0
//     for end := 0; end < n; end++ {
//         currentChar := s[end]
//         if lastIndex[currentChar] > start {
//             start = lastIndex[currentChar]
//         }
//         if end-start+1 > maxLength {
//             maxLength = end - start + 1
//         }

//         lastIndex[currentChar] = end + 1

//     }

//     return maxLength
// }

func main() {
	s := "dvdf"

	// fmt.Printf("%c", s[0])
	fmt.Println(lengthOfLongestSubstring(s))
}
