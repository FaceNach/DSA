package main

import "fmt"

func bubbleSort(sort []int) []int {

	for i := 0; i < len(sort); i++ {
		for j := 0; j < len(sort)-1; j++ {
			if sort[j] > sort[j+1] {
				sort[j], sort[j+1] = sort[j+1], sort[j]
			}
		}
	}

	return sort
}

func main() {

	sort := []int{1, 3, 2, 9, 12, 8, 4}

	sorted := bubbleSort(sort)

	for _, val := range sorted {
		fmt.Println(val)
	}

}
