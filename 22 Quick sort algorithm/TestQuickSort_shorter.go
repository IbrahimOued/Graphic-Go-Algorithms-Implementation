package main

import "fmt"

func main() {

	var scores = []int{50, 65, 99, 87, 74, 63, 76, 100}
	var sortedScores = quickSort(scores)
	fmt.Printf("%d", sortedScores)
}

func quickSort(data []int) []int {
	if len(data) <= 1 {
		return nil
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			// when data[i] == mid, an unstable sort takes place
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	quickSort(data[:head])
	quickSort(data[head+1:])
	return data
}
