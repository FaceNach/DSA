package main

import "fmt"

// Frequency Counter - constructNote
// Write a function called constructNote, which accepts two strings, a message and some letters.
// The function should return true if the message can be built with the letters that you are given, or it should return false.
// Assume that there are only lowercase letters and no space or special characters in both the message and the letters.


func constructNote(str1, str2 string) bool{
	if len(str1) > len(str2){
		return false
	}
	
	str1Map := make(map[rune]int)
	str2Map := make(map[rune]int)
	
	for _, val := range str1 {
		str1Map[val]++
	}
	
	for _, val := range str2 {
		str2Map[val]++
	}
	
	for key, val := range str1Map{
		if val > str2Map[key]{
			return false
		}
	}
	
	return true
}

func constructNoteOptimized(msg, letters string) bool{
	if len(msg) > len(letters){
		return false
	}
	
	lettersMap := make(map[rune]int)
	
	for _, val := range letters{
		lettersMap[val]++
	}
	
	for _, val := range msg{
		if lettersMap[val] == 0 {
			return false
		}
		
		lettersMap[val]--
	}
	
	return true
}

func main() {
	fmt.Println("--- Versión Normal (2 Mapas) ---")
    fmt.Println(constructNote("aa", "aab")) 
    
    fmt.Println(constructNote("aa", "abc")) 
    

    fmt.Println(constructNote("abc", "abcd")) 


    fmt.Println("\n--- Versión Optimizada (1 Mapa) ---")
    fmt.Println(constructNoteOptimized("aa", "aab")) 
    
    fmt.Println(constructNoteOptimized("aa", "abc")) 
    
    fmt.Println(constructNoteOptimized("holamundo", "hola"))
}