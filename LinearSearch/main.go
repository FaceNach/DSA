package main

import "fmt"

func linearSearch(arr []int, v int) (int, bool) {
	for i, val := range arr {
		if val == v {
			return i, true
		}
	}

	return -1, false
}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	value := 7

	index, find := linearSearch(arr, value)

	if find {

		fmt.Printf("You found the value %v on the index %v", value, index)
	} else {
		fmt.Println("The value it's not in the array ")
	}

}
