package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Please enter the index to be deleted : \n")
	var index int
	fmt.Scan(&index)

	var length = len(scores)
	var tempArray = make([]int, length-1) // Create a new array
	for i := 0; i < length; i++ {
		if i < index { // Copy the array after index to the end of tempArray
			tempArray[i] = scores[i]
		}
		if i > index { // Copy the array after indx to the end of tempArray
			tempArray[i-1] = scores[i]
		}
	}
	scores = tempArray
	for i := 0; i < length-1; i++ {
		fmt.Printf("%d, ", scores[i])
	}
}
