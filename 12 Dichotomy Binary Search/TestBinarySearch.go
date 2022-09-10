package main

import "fmt"

func main() {
	var scores = []int{30, 40, 50, 70, 85, 90, 100}
	var searchValue = 40
	var position = binarySearch(scores, len(scores), searchValue)
	fmt.Printf("%d position : %d", searchValue, position)
	fmt.Printf("\n--------------------------\n")

	searchValue = 90
	position = binarySearch(scores, len(scores), searchValue)
	fmt.Printf("%d position : %d", searchValue, position)
	fmt.Printf("\n--------------------------\n")
}

func binarySearch(arrays []int, length int, searchValue int) int {
	var low = 0
	var high = length
	var mid = 0

	// This is a while loop in Go
	for {
		if low >= high {
			break
		}
		mid = (low + high) / 2
		if arrays[mid] == searchValue {
			return mid
		} else if arrays[mid] < searchValue {
			low = mid + 1
		} else if arrays[mid] > searchValue {
			high = mid - 1
		}
	}
	return -1
}
