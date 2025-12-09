package main

import (
	"fmt"
)

func maxSubarraySum( arr []int, num int) int {
	if len(arr) == 0{
		return 0
	}
	
	if num == 0 {
		return 0
	}
	
	if num > len(arr){
		return 0
	}
	
	var maxCount int
	
	for i:=0; i < len(arr) - num + 1; i++ {
		
		count := 0
		
		for j:=0; j< num; j++ {
			count += arr[i +j ]
		}
		
		if(count > maxCount){
			maxCount = count
		}
	}
	return maxCount;
} 

func maxSubarraySumOptimized( arr []int, num int) int {
	maxSum :=0
	tempSum :=0
	
	if num > len(arr){
		return 0
	}
	
	for i:=0; i< num; i++ {
		maxSum +=arr[i]
	}
	
	tempSum = maxSum;
	
	for i:= num; i < len(arr); i++ {
		tempSum = tempSum - arr[i - num] + arr[i]
		if tempSum > maxSum{
			maxSum = tempSum
		}
	}
	
	return maxSum;
	
}

func main() {
	
	arr1 := []int {1,2,5,2,8,1,5}
	res1 := maxSubarraySum(arr1,2)
	
	res2 := maxSubarraySumOptimized(arr1,2)
	
	fmt.Println(res1)
	fmt.Println(res2)
	
}