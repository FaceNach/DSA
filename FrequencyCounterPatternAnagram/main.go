package main

import "fmt"

func validAnagram(str1, str2 string) bool{
	
	if len(str1) != len(str2){
		return false
	}
	
	freq1 := make(map[rune]int)

	for _, val := range str1 {
		freq1[val]++
	}
	
	for _, val := range str2{
		
		
		//another way to apply this frequency count pettern
		// count, ok := freq1[val]
		
		// if !ok {
		// 	return false
		// }
		
		// if count == 0{
		// 	return false
		// }
		
		if freq1[val] == 0{
			return false
		}
		
	 	freq1[val]--
	}
	return true
}

func main(){
	fmt.Println("--- Casos Básicos ---")
		fmt.Printf("1. 'anagram' vs 'nagaram': %v\n", validAnagram("anagram", "nagaram"))
		fmt.Printf("2. 'rat' vs 'car':         %v\n", validAnagram("rat", "car"))
		fmt.Printf("3. 'gato' vs 'gatos':      %v\n", validAnagram("gato", "gatos"))
		fmt.Printf("4. 'aaz' vs 'zza':         %v\n", validAnagram("aaz", "zza"))

		fmt.Println("\n--- Casos Especiales (Runes) ---")
		fmt.Printf("5. 'amor' vs 'roma':       %v\n", validAnagram("amor", "roma"))
		fmt.Printf("6. 'niño' vs 'noñi':       %v\n", validAnagram("niño", "noñi"))
		fmt.Printf("7. 'hola👋' vs '👋hola':    %v\n", validAnagram("hola👋", "👋hola"))
		fmt.Printf("8. '' vs '':               %v\n", validAnagram("", ""))
}