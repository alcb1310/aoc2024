package main

import "math"

func distance(left []int, right []int) int {
	left = sortArray(left)
	right = sortArray(right)

	dist := 0
	for i := range left {
		dist += int(math.Abs(float64(right[i] - left[i])))
	}

	return dist
}

func sortArray(x []int) []int {
  
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
