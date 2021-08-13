package package_dua

// buat variabel primes dan buatlah
// menggunakan array literals dengan panjang 6
// buat 6 angka bilangan prima
// print primes

func Primes() [6]int {
	var primes [6]int
	val := 0
	i := 2
	x := 0
	for val < 6 {
		if i%2 == 1 {
			primes[x] = i
			x++
			val++
		}
		i++
	}
	return primes
}
