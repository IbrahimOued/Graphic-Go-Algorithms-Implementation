package main

import "fmt"

func main() {
	var scores = []int{60, 80, 95, 50, 75}

	fmt.Print(selection_sort(scores, len(scores)))
}

func selection_sort(arrays []int, length int) []int {
	var minIndex int // Save the index of the selected minimum

	for i := 0; i < length; i++ {
		minIndex = i
		// Save the minimum value of each loop as the 1st element
		var minValue = arrays[minIndex]
		for j := i; j < length-1; j++ {
			if minValue > arrays[j+1] {
				minValue = arrays[j+1]
				minIndex = j + 1
			}
		}

		// If minIndex changed, current minimum is exchanged with minIndex
		if i != minIndex {
			var temp = arrays[i]
			arrays[i] = arrays[minIndex]
			arrays[minIndex] = temp
		}
	}
	return arrays
}
