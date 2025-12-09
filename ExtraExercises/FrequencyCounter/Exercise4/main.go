package main

// Frequency Counter - findAllDuplicates
// Given an array of positive integers, some elements appear twice and others appear once. 
// Find all the elements that appear twice in this array. Note that you can return the elements in any order.


import "sort"

func findDuplicates(nums ...int ) []int{
	if len(nums) == 0 {
		return []int{}
	}
	
	sort.Ints(nums)
	
	start := 0
	next := 1
	
	out := make([]int,0)
	
	for next < len(nums){
		if nums[start] == nums[next]{
			out = append(out, nums[start])
		}
		
		start++
		next++
	}
	
	return out
}

func main() {
	
}