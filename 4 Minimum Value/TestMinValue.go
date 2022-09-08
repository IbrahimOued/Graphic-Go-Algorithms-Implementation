package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex = 0 // The index of the minimum
	for j := 1; j < length; j++ {
		if arrays[minIndex] > arrays[j] {
			minIndex = j
		}
	}
	return arrays[minIndex]
}

func main() {
	var scores = []int{60, 80, 95, 50, 75}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Printf("Min value = %d\n", minValue)
}
