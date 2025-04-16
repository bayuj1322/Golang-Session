package session3

import "fmt"

type isOddNum func(int) bool

func findOddNumbers(numbers []int, callback isOddNum) int {
	totalOddNumbers := 0
	for _, num := range numbers {
		if callback(num) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

func Sesi31() {
	// Slice berisi angka-angka
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Kirim closure sebagai callback yang cek apakah angka ganjil
	result := findOddNumbers(numbers, func(n int) bool {
		return n%2 != 0
	})

	fmt.Println("Total angka ganjil:", result)
}
