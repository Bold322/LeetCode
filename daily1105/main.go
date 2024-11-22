package main

import "fmt"

func minChanges(s string) int {
	count := 0
	runes := []rune(s)
	for i := 0; i < len(s)-1; i = i + 2 {
		if s[i] != s[i+1] {
			if runes[i] == '1' {
			    runes[i+1] = '1'
			} else {
			    runes[i+1] = '0'
			}
			count++
		}
	}

	fmt.Println(string(runes))
	return count
}

func main() {
	s := "10010101010101111110001001011010101"

	fmt.Println(minChanges(s))
}
