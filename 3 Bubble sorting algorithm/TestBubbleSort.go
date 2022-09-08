package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	bubble_sort(scores, len(scores))
	fmt.Print(bubble_sort(scores, len(scores)))

}

func bubble_sort(arrays []int, length int) []int {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arrays[j] > arrays[j+1] { // Swap
				var flag = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag
			}
		}
	}
	return arrays
}
