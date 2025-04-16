package session3

import "fmt"

func changeValue(x *int) {
	*x = 20 // mengubah nilai dari variabel yang ditunjuk oleh pointer
}

func Sesi32() {
	a := 10
	fmt.Println("Sebelum:", a)

	changeValue(&a) // mengirim alamat memori dari a

	fmt.Println("Sesudah:", a)

}
