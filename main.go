package main

import (
	"fmt"
	"sort"
)

//problem-1
// Input:  []int{3, 1, 4, 1, 5, 9, 2, 6}
// Output: evens = [4 2 6], odds = [3 1 1 5 9]

// problem-2
// Input:  []string{"error", "info", "error", "debug", "error", "info"}
// Output: map[debug:1 error:3 info:2]
func main() {
	input := []string{"error", "info", "error", "debug", "error", "info"}

	result := problem2(input)

	keys := make([]string, 0, len(result))
	for k, v := range result {
		if v > 1 {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, result[k])
	}
}

func problem2(n []string) map[string]int {

	r := make(map[string]int)

	for _, v := range n {
		r[v]++
	}
	return r
}

func problem1(n []int) ([]int, []int) {
	even := make([]int, 0, len(n))
	odd := make([]int, 0, len(n))

	for _, v := range n {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return even, odd
}
