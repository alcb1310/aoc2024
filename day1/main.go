package main

import (
	"fmt"
	"math"
)

func main() {
	_ = sample
	_ = data

	var left []int
	var right []int
	for _, line := range data {
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, a)
		right = append(right, b)
	}

	left = sort(left)
	right = sort(right)

	dist := 0
	for i := range left {
		dist += int(math.Abs(float64(right[i] - left[i])))
	}

	fmt.Println(dist)
}

func sort(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for y := 1; y < len(x); y++ {
			if x[y-1] > x[y] {
				a := x[y-1]
				x[y-1] = x[y]
				x[y] = a
			}
		}
	}
	return x
}
