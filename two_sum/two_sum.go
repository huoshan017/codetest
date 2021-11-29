package main

import (
	"log"
)

func two_sum(sum int, array []int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(array); i++ {
		half := sum - array[i]
		if idx, o := m[half]; o {
			return i, idx
		}
		m[array[i]] = i
	}
	return -1, -1
}

func main() {
	sum := 20
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	idx1, idx2 := two_sum(sum, array[:])
	log.Printf("the two index %v %v", idx1, idx2)
}
