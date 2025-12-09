package main

func countUniqueValues(arr []int) int{
	if (len(arr) == 0){
		return 0
	}
	
	if len(arr) == 1{
		return 1
	}
	
	uniqueVals := make(map[int]int)
	
	for _, val := range arr {
		uniqueVals[val]++
	}
	
	return len(uniqueVals)
}

func countUniqueValuesWithTwoPointerMethod(arr []int) int {
	if len(arr) == 0{
		return 0
	}
	
	i:= 0
	
	for j:=1 ; j< len(arr); j++{
		if arr[i] != arr[j]{
			i++
			arr[i] = arr[j]
		}
	}
	
	return i + 1
}


func main() {
	
	
}