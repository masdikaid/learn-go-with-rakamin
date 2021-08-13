package package_satu

import "fmt"

// buat array a string dengan panjang 2
// masukkan nilai pertama "Go"
// masukkan nilai kedua "C"
// print nilai pertama, dan kedua
// print a

func Arraya() {
	array := [2]string{"Go", "C"}
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array)
}
