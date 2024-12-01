package main

func similarity(left []int, right []int) int {
	total := 0
	for _, v := range left {
		if val := search(v, right); val > 0 {
      x := val * v
      total += x
		}
	}
	return total
}

func search(value int, x []int) int {
	appear := 0
	for _, v := range x {
		if v == value {
			appear++
		}
	}
	return appear
}
