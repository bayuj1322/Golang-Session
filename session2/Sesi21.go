package session2

import "fmt"

func Sesi21() {
	tinggi := 5
	for i := 1; i <= tinggi; i++ { // loop baris
		fmt.Print(i, ". ")
		for j := 1; j <= tinggi-i; j++ { // loop kolom
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print(k, " ")
		}
		fmt.Println() // pindah baris
	}

}
