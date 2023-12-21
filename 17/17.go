package main

import "fmt"

func main() {
	sequence := []int{1, 4, 10, 11, 18, 24, 28, 50}
	numToSearch := 4
	fmt.Println(binarySearch(sequence, numToSearch))
}

func binarySearch(sortedSequence []int, value int) bool {
	lastSplit := len(sortedSequence) / 2
	for {
		if sortedSequence[lastSplit] > value {
			sortedSequence = sortedSequence[:lastSplit]
			lastSplit /= 2
		} else if sortedSequence[lastSplit] == value {
			return true
		} else if sortedSequence[lastSplit] < value {
			sortedSequence = sortedSequence[lastSplit:]
			lastSplit /= 2
		}
		if lastSplit == 0 {
			return false
		}
	}
}
