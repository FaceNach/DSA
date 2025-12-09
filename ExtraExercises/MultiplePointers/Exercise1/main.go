package main

import "fmt"

// Multiple Pointers - averagePair
// Write a function called averagePair.
// Given a sorted array of integers and a target average, determine if there is a pair of values in the array where the average of the pair equals the target average.
// There may be more than one pair that matches the average target.


func averagePair(target float32,nums ...int) bool{
	if len(nums) == 0{
		return false
	}
	
	start := 0
	end := len(nums) -1
	
	for start < end{
		
		avg := float32((nums[start] + nums[end])) / 2
		
		if   avg == target {
			return true
		}
		
		if avg > target{
			end--
		}else{
			start++
		}
		
	}
	
	
	return false
}

func main() {
	
    fmt.Println(averagePair(2.5, 1, 2, 3)) 

    fmt.Println(averagePair(8, 1, 2, 6, 7, 8, 12, 21))


    fmt.Println(averagePair(4, 1, 3, 3, 5, 6, 7, 10, 12, 19)) // Returns true actually! (1+7)/2 = 4.


    fmt.Println(averagePair(1)) 

    fmt.Println(averagePair(10, 1, 2, 3))
}