package main

import "fmt"

func main() {
	var scores = []int{50, 60, 70, 80, 90}
	var length = len(scores)

	fmt.Print(reverse(scores, length))
}

func reverse(arrays []int, length int) []int {
	var middle = length / 2
	for i := 0; i <= middle; i++ {
		// var temp = arrays[i]
		// arrays[i] = arrays[lenght-i-1]
		// arrays[lenght-i-1] = temp
		arrays[i], arrays[length-i-1] = arrays[length-i-1], arrays[i]
	}
	return arrays
}
