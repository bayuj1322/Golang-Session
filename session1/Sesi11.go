package session1

import "fmt"

func Sesi11() {
	// Operator Aritmatika
	a := 10
	b := 3

	fmt.Println("== Operator Aritmatika ==")
	fmt.Println("a + b =", a+b) // Penjumlahan
	fmt.Println("a - b =", a-b) // Pengurangan
	fmt.Println("a * b =", a*b) // Perkalian
	fmt.Println("a / b =", a/b) // Pembagian (integer division)
	fmt.Println("a % b =", a%b) // Modulus (sisa bagi)

	// Operator Perbandingan (Relasional)
	fmt.Println("\n== Operator Perbandingan ==")
	fmt.Println("a == b =", a == b) // Sama dengan
	fmt.Println("a != b =", a != b) // Tidak sama dengan
	fmt.Println("a > b =", a > b)   // Lebih dari
	fmt.Println("a < b =", a < b)   // Kurang dari
	fmt.Println("a >= b =", a >= b) // Lebih dari sama dengan
	fmt.Println("a <= b =", a <= b) // Kurang dari sama dengan

	// Operator Logika (Boolean)
	x := true
	y := false

	fmt.Println("\n== Operator Logika ==")
	fmt.Println("x && y =", x && y) // AND
	fmt.Println("x || y =", x || y) // OR
	fmt.Println("!y =", !y)         // NOT

}
