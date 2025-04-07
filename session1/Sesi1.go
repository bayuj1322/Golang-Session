package session1

import "fmt"

func Sesi1() {
	// Tipe Data Angka (Number)
	var angkaInt int = 10               // integer
	var angkaInt8 int8 = -128           // integer 8-bit
	var angkaUint uint = 100            // unsigned integer
	var angkaFloat float64 = 3.14       // float 64-bit
	var angkaComplex complex64 = 2 + 3i // bilangan kompleks

	// Tipe Data Boolean
	var benar bool = true
	var salah bool = false

	// Tipe Data String
	var nama string = "Golang"
	var karakter byte = 'A'     // alias uint8
	var karakterRune rune = '❤' // alias int32

	// Menampilkan Semua
	fmt.Println("Integer:", angkaInt, angkaInt8, angkaUint)
	fmt.Println("Float:", angkaFloat)
	fmt.Println("Complex:", angkaComplex)
	fmt.Println("Boolean:", benar, salah)
	fmt.Println("String:", nama, string(karakter), string(karakterRune))

}
