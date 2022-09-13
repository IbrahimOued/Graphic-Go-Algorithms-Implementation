package main

import "fmt"

func main() {
	var scores = []int{50, 65, 99, 87, 74, 63, 76, 100}
	var sortedScores = mergeSort(scores)
	fmt.Printf("%d", sortedScores)
}

// mergeSort Algorithm
// Time complexity O(n*logN)
// Space complexity O(n)
func mergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}

	key := n / 2
	left := mergeSort(array[0:key])
	right := mergeSort(array[key:])
	return merge(left, right)
}

// Merge 2 sorted array
func merge(left []int, right []int) []int {
	tmp := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	// Eahc of the array is traversed
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	// A temp array is used to backed the result
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
