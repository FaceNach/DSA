package main

import (
	"fmt"
	"slices"
)

func same(arr1, arr2 []int) bool{
	if len(arr1) != len(arr2){
		return false
	}
	
	for _, val1 := range arr1{
	 	idx := slices.Index(arr2, val1 * val1)
			
		if idx == -1{
			return false
		}
		arr2 = append(arr2[:idx], arr2[:idx+1]... )
	}
	
	return true
}

func same2(arr1, arr2 []int) bool{
	if len(arr1) != len(arr2){
		return false
	}
	
	for  i := 0; i < len(arr1) ; i++ {
		idx := -1
		
		for j:=0; j < len(arr2); j++ {
			if arr1[i] * arr1[i] == arr2[j]{
				idx = j
				break
			}
		}
		
		if idx == -1{
			return false
		}
		
		arr2 = append(arr2[:idx],arr2[idx+1:]... )
	}
	
	return true
}

func sameOptimized(arr1, arr2 []int) bool{
	if (len(arr1) != len(arr2)){
		return false
	}
	
	freq := make(map[int]int)
	
	for _, val := range arr2 {
		freq[val]++
	}
	
	for _, val := range arr1{
		squared := val * val
		if count, exist := freq[squared]; !exist || count == 0 {
			return false
		}
		freq[squared]--
	}
	
	return true
}


func main(){
	
	arr1 := []int {1,2,3}
	arr2 := []int {1,4,9}
	arr3 := []int {5,2,3}
	arr4 := []int {1,4,9}
	
	rsl := same(arr1, arr2)
	rsl2 := same(arr3, arr4)
	
	fmt.Printf("%v \n", rsl)
	fmt.Printf("%v", rsl2)
	
}