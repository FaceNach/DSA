package main

import "fmt"

//assumption that the array its always sorted, if its not the case you cannot binary search it!

func binarySearch(arr []int, v int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := low + (high-low)/2

		if arr[mid] == v {
			return true
		} else if v > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	value := 10

	result := binarySearch(arr, value)

	if result {
		fmt.Printf("The value %v its on the array", value)
	} else {
		fmt.Printf("The value %v its not in the array", value)
	}

}
