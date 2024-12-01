package main

import (
	"fmt"
)

func main() {
	_ = sample
	_ = data

	var left []int
	var right []int
	for _, line := range data {
		// for _, line := range sample {
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, a)
		right = append(right, b)
	}

	fmt.Println("Distance: ", distance(left, right))
	fmt.Println("Similarity: ", similarity(left, right))
}
