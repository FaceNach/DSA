package main

import "fmt"

func isSubsequence(str1, str2 string) bool {
	if len(str1) == 0 || len(str2) == 0{
		return false
	}
	
	j:=0
	i:=0
	
	for j < len(str2){
		if(str1[i] == str2[j]){
			i++
		}
		
		if i == len(str1){
			return true
		}
		
		j++
	}
	
	return false
}

func main() {

    fmt.Println(isSubsequence("sing", "sting")) 

    fmt.Println(isSubsequence("abc", "abracadabra")) 

    fmt.Println(isSubsequence("acb", "abc")) 


    fmt.Println(isSubsequence("xyz", "x y z"))

    
    fmt.Println(isSubsequence("abc", "ac"))
}