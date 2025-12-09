package main

import (
	"fmt"
)

//the array HAS TO BE  sorted...
func findIndexSortedArray(arr []int, num int ) int{
	
	var Zero int
	
	 if len(arr) == 0{
			return Zero
		}
	
	min := 0
	max := len(arr) - 1;
	
	
	for min <= max {
		middle := int((min + max)/2)
		currentElement := arr[middle]
		
		if currentElement == num {
			return middle
		}
		
		if currentElement < min {
			min = middle +1
		}else{
			max = middle - 1
		}
	}
	
	return -1
}

func main() {
	
	sortedList := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23
	result := findIndexSortedArray(sortedList, target)
	fmt.Printf("Searching for %d: Found at index %d\n", target, result)

	targetMissing := 10
	resultMissing := findIndexSortedArray(sortedList, targetMissing)
	fmt.Printf("Searching for %d: Found at index %d (Note: -1 means not found)\n", targetMissing, resultMissing)

	targetFirst := 2
	resultFirst := findIndexSortedArray(sortedList, targetFirst)
	fmt.Printf("Searching for %d: Found at index %d\n", targetFirst, resultFirst)

	emptyList := []int{}
	resultEmpty := findIndexSortedArray(emptyList, 5)
	fmt.Printf("Searching in empty array: Found at index %d\n", resultEmpty)
	
}