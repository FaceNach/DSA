package main

import (
	"cmp"
	"fmt"
	"sort"
)

// Frequency Counter / Multiple Pointers - areThereDuplicates
// Implement a function called, areThereDuplicates which accepts a variable number of arguments, and checks whether there are any duplicates among the arguments passed in.
// You can solve this using the frequency counter pattern OR the multiple pointers pattern.

func areThereDuplicates[T comparable](arr ...T) bool{
	if(len(arr) == 0 || len(arr) == 1){
		return false
	}
	
	for i := 0; i < len(arr); i++ {
		for j := i +1 ; j < len(arr); j++{
			if arr[i] == arr[j]{
				return true
			}
		}
	}
	
	return false
}

func areThereDucplicatesOptimized[T comparable] (arr ...T) bool{
	if(len(arr) == 0 || len(arr) == 1){
		return false
	}
	
	check := make(map[T]int)
	
	for _, val := range arr{
		check[val]++
		if check[val] > 1{
			return true
		}
	}
	
	return false
}

func areThereDuplicatesTwoPointersAproach[T cmp.Ordered](arr ...T) bool {
	sort.Slice(arr, func(i, j int) bool{
		return arr[i] < arr[j]
	})
	
	
	start :=0
	next := 1
	
	for next < len(arr){
		if arr[start] == arr[next]{
			return  true
		}
		
		start++
		next++
	}
	
	return false
}

func main() {
	
	fmt.Println("--- 1. Testing Brute Force (Nested Loops) ---")
		fmt.Println("1, 2, 3:", areThereDuplicates(1, 2, 3))             // false
		fmt.Println("1, 2, 2:", areThereDuplicates(1, 2, 2))             // true
		fmt.Println("a, b, c, a:", areThereDuplicates("a", "b", "c", "a")) // true

		fmt.Println("\n--- 2. Testing Frequency Counter (Map) ---")
		fmt.Println("1, 2, 3:", areThereDucplicatesOptimized(1, 2, 3))       // false
		fmt.Println("1, 2, 2:", areThereDucplicatesOptimized(1, 2, 2))       // true
		fmt.Println("apple, banana, apple:", areThereDucplicatesOptimized("apple", "banana", "apple")) // true

		fmt.Println("\n--- 3. Testing Two Pointers (Sort) ---")
		
		fmt.Println("1, 2, 3:", areThereDuplicatesTwoPointersAproach(1, 2, 3))             // false
		fmt.Println("1, 2, 2:", areThereDuplicatesTwoPointersAproach(1, 2, 2))             // true
		fmt.Println("x, y, z, x:", areThereDuplicatesTwoPointersAproach("x", "y", "z", "x")) // true
		
		fmt.Println("1.5, 2.5, 1.5:", areThereDuplicatesTwoPointersAproach(1.5, 2.5, 1.5)) // true

		fmt.Println("\n--- 4. Watch out for Side Effects! ---")
		myData := []int{5, 1, 3, 2}
		fmt.Println("Original Data:", myData)
		
		areThereDuplicatesTwoPointersAproach(myData...)
		
		fmt.Println("Data After TwoPointers (It got sorted!):", myData)
	
}