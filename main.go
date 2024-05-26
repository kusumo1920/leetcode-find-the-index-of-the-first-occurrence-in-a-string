package main

import "fmt"

func main() {
	haystack := "leetcode"
	needle := "leeto"
	output := strStr(haystack, needle)
	fmt.Println(output)
}

func strStr(haystack, needle string) int {
	return solution1(haystack, needle, 0)
}

func solution1(haystack, needle string, index int) int {
	if index+len(needle) > len(haystack) {
		return -1
	}

	if needle == haystack[index:index+len(needle)] {
		return index
	}

	index++
	return solution1(haystack, needle, index)
}
