package main

import (
	"fmt"
	"math"
)

func twoCrystalBalls(a []bool) int {

	jumpAmount := int(math.Sqrt(float64(len(a))))

	i := jumpAmount

	for ; i < len(a); i += jumpAmount {
		if a[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(a); j, i = j+1, i+1 {
		if a[i] {
			return i
		}
	}

	return -1
}

func main() {

	pass := []bool{false, false, false, false, false, true, true, true, true}
	pass2 := []bool{false, false, false, false, false, false, false, false, false}

	result := twoCrystalBalls(pass)
	result2 := twoCrystalBalls(pass2)

	fmt.Println(result)
	fmt.Println(result2)

}
