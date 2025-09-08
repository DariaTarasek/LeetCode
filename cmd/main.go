package main

import (
	"LeetCodeSolutions/maps_frequencies"
	_ "LeetCodeSolutions/strings_slices"
	"fmt"
)

func main() {
	fmt.Println(maps_frequencies.TopKFrequent([]int{1, 1, 0, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}, 2))
}
