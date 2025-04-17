package session4

import "fmt"

func Sesi41() {
	// rs adalah slice of empty interface
	rs := []interface{}{"teks", 123, true, 3.14}

	// rm adalah map dengan string sebagai key dan empty interface sebagai value
	rm := map[string]interface{}{
		"nama":    "Budi",
		"umur":    30,
		"menikah": false,
		"tinggi":  170.5,
	}

	// Menampilkan isi dari slice rs
	fmt.Println("Isi slice rs:")
	for i, v := range rs {
		fmt.Printf("Index %d: %v (tipe: %T)\n", i, v, v)
	}

	// Menampilkan isi dari map rm
	fmt.Println("\nIsi map rm:")
	for k, v := range rm {
		fmt.Printf("Key '%s': %v (tipe: %T)\n", k, v, v)
	}
}
