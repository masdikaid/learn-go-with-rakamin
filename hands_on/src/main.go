package main

import (
	"fmt"
	"hands_on/package_dua"
	"hands_on/package_empat"
	"hands_on/package_satu"
	"hands_on/package_tiga"
)

func main() {

	// hands on 1
	// buat array a string dengan panjang 2
	// masukkan nilai pertama "Go"
	// masukkan nilai kedua "C"
	// print nilai pertama, dan kedua
	// print a
	// buat variabel primes dan buatlah
	// menggunakan array literals dengan panjang 6
	// buat 6 angka bilangan prima
	// print primes
	fmt.Println("")
	fmt.Println("===== tugas 1 =====")
	package_satu.Arraya()
	fmt.Println("===== ======= =====")
	fmt.Println("")
	fmt.Println("===== tugas 2 =====")
	fmt.Println(package_dua.Primes())
	fmt.Println("===== ======= =====")
	fmt.Println("")

	// Hands on 2
	// buat variabel primes dan buatlah
	// menggunakan slice literals dengan panjang 6
	// buat 6 angka bilangan prima
	// print primes
	// print primes dari data ke 3 sampai 5
	fmt.Println("")
	fmt.Println("===== tugas 1 =====")
	primes := package_dua.Primes()
	fmt.Println(primes[2:5])
	fmt.Println("===== ======= =====")
	fmt.Println("")

	// hands on 3
	// Menggunakan Make
	// Buat slice of int bernama "a" dengan panjang 5
	// Buat slice of int bernama "b" dengan panjang 0 kapasitas
	// Buat variabel c dengan dua data awal b
	// Buat variabel d dengan data ke 2 sampai 4
	fmt.Println("")
	fmt.Println("===== tugas 1 =====")
	a := package_tiga.SliceBuilder(5)
	package_tiga.PrintSlice("a", a)
	fmt.Println("===== ======= =====")
	fmt.Println("")
	fmt.Println("===== tugas 2 =====")
	b := package_tiga.SliceBuilder(10)
	package_tiga.PrintSlice("b", b)
	fmt.Println("===== ======= =====")
	fmt.Println("")
	fmt.Println("===== tugas 3 =====")
	c := b[:2]
	package_tiga.PrintSlice("c", c)
	fmt.Println("===== ======= =====")
	fmt.Println("")
	fmt.Println("===== tugas 4 =====")
	d := b[1:4]
	package_tiga.PrintSlice("d", d)
	fmt.Println("===== ======= =====")
	fmt.Println("")

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
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("")
	fmt.Println("===== tugas 1 =====")
	package_empat.StringBuilder(pow)
	fmt.Println("===== ======= =====")
	fmt.Println("")

	// hands on 5
	// buat variabel m dengan map key: string dan value: int
	// beri nilai "Answer" dengan 42
	// ganti nilai "Answer" dengan 48
	// hapus "Answer
	// gunakan pengecekan: elem, ok = m[key]
	var m map[string]int
	m = map[string]int{"answer": 42}
	m["answer"] = 48
	delete(m, "answer")
	fmt.Println("")
	fmt.Println("===== tugas 1 =====")
	elem, ok := m["answer"]
	fmt.Println(elem, ok)
	fmt.Println("===== ======= =====")
	fmt.Println("")
}
