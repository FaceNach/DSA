package main

import (
	"fmt"
	"strconv"
)

// Write a function called sameFrequency. Given two positive integers, find out if the two numbers have the same frequency of digits.
// Your solution MUST have the following complexities:
// Time: O(N)

func sameFrequency(num1, num2 int) bool{
	
	str1 := strconv.Itoa(num1)
	str2 := strconv.Itoa(num2)
	
	if len(str1) != len(str2){
		return false
	}
	
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)
	
	for _, num := range str1{
		map1[num]++
	}
	
	for _, num := range str2{
		map2[num]++
	}
	
	if len(map1) != len(map2) {
        return false
    }
	
	for key, value := range map1{
		if value != map2[key]{
			return false
		}
	}
	return true
}


	func main(){
	fmt.Println(sameFrequency(182, 281))
	fmt.Println(sameFrequency(34, 14))
	fmt.Println(sameFrequency(3589578, 5879385))
	fmt.Println(sameFrequency(22, 222))
	}