package package_empat

import (
	"fmt"
)

// handson 4
// buatlah for dengan range untuk pow
// yang menghasilkan:
/*
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
*/

func StringBuilder(arr []int) {
	for _, value := range arr {
		fmt.Printf("2**%d = %d\n", findPower(value), value)
	}
}

func findPower(x int) int {
	n := x
	val := 0
	for n/2 > 0 {
		n = n / 2
		val++
	}
	return val
}
