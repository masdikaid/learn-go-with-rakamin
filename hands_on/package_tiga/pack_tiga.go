package package_tiga

import "fmt"

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SliceBuilder(length int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = i
	}
	return slice
}
