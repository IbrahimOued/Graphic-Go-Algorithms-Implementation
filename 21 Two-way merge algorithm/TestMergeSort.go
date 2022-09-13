package main

import "fmt"

func main() {
	var scores = []int{50, 65, 99, 87, 74, 63, 76, 100}
	var length = len(scores)
	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(array []int, length int) {
	var temp = make([]int, length)
	mergeSort(array, temp, 0, length-1)
}

func mergeSort(array []int, temp []int, left int, right int) {
	if left < right {
		var center = (left + right) / 2
		mergeSort(array, temp, left, center)      // Left merge sort
		mergeSort(array, temp, center+1, right)   // Right merge sort
		merge(array, temp, left, center+1, right) // Merge two ordered arrays
	}
}

/**
Combine two ordered list into an ordered list
temp : Temporary array
left : Start the subscript on the left
right : Start the subscript on the right
rightEndIndex : End subscript on the right
*/

func merge(array []int, temp []int, left int, right int, rightEndIndex int) {
	var leftEndIndex = right - 1 // End subscript on the left
	var tempIndex = left         // Starting from the left count
	var elementNumber = rightEndIndex - left + 1
	for {
		if left > leftEndIndex || right > rightEndIndex {
			break
		}
		if array[left] <= array[right] {
			temp[tempIndex] = array[left]
			tempIndex++
			left++
		} else {
			temp[tempIndex] = array[right]
			tempIndex++
			right++
		}
	}
	for {
		if left > leftEndIndex {
			break
		}
		temp[tempIndex] = array[left] // If there is element on the left
		tempIndex++
		left++
	}
	for {
		if right > rightEndIndex {
			break
		}
		temp[tempIndex] = array[right] // If there is element on the right
		tempIndex++
		right++
	}
	for i := 0; i < elementNumber; i++ {
		array[rightEndIndex] = temp[rightEndIndex]
		rightEndIndex--
	}
}
