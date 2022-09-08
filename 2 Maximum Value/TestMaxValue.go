package main

import "fmt"

func Max(arrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] { // Make the swap
			var temp = arrays[i]
			arrays[i+1] = arrays[i]
			arrays[i+1] = temp
		}

	}
	var maxValue = arrays[length-1]
	return maxValue
}

func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var maxValue = Max(scores, length)
	fmt.Printf("Max %d\n", maxValue)
}
