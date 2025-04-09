package session2

import "fmt"

func Sesi22() {
	// Array 3 Dimensi [2][2][3] short declaration
	data := [2][2][3]int{
		{
			{10, 20, 30},
			{40, 50, 60},
		},
		{
			{70, 80, 90},
			{100, 110, 120},
		},
	}

	fmt.Println("Isi Array 3 Dimensi:")

	for i, lantai := range data { // i itu index lantai
		fmt.Println("Lantai ke-", i+1)

		for j, baris := range lantai { // j itu index baris
			for k, angka := range baris { // k itu index kolom, angka itu valuenya langsung
				fmt.Printf("Data di [Lantai %d][Baris %d][Kolom %d] = %d\n", i+1, j+1, k+1, angka)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
