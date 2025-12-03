package main

import "fmt"

//we asummed that this is a sorted array
func sumZero(sortArr []int) ([]int){
	if (len(sortArr) == 0 || len(sortArr) == 1){
		return []int{}
	}
	
	left := 0
	right := len(sortArr) -1
	
	for left < right{
		sum := sortArr[left] + sortArr[right]
		
		if (sum == 0){
			return []int{sortArr[left], sortArr[right]}
		}
		
		if sum > 0 {
			right--
		}
		
		if sum < 0 {
			left++
		}
	}
	
	return []int{}
}

func main() {
	
	fmt.Println(sumZero([]int{-3, -2, -1, 0, 1, 2, 3})) // [-3, 3]
	fmt.Println(sumZero([]int{-2, 0, 1, 3}))            // []
	fmt.Println(sumZero([]int{1, 2, 3}))                // []
	fmt.Println(sumZero([]int{-4, -3, -2, 1, 2, 5}))    // [-2, 2]
	
}