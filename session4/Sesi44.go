package session4

import (
	"fmt"
	"sync"
)

// Fungsi untuk mencetak nama buah
func printFruit(fruit string, index int, wg *sync.WaitGroup) {
	defer wg.Done() // Memberi tahu bahwa goroutine ini sudah selesai

	fmt.Printf("Buah ke-%d: %s\n", index+1, fruit)
}

func Sesi44() {
	// Daftar buah
	fruits := []string{"Apel", "Jeruk", "Mangga", "Pisang"}

	// Buat instance WaitGroup
	var wg sync.WaitGroup

	// Loop setiap buah, jalankan printFruit sebagai goroutine
	for i, fruit := range fruits {
		wg.Add(1) // Tambah counter goroutine yang harus ditunggu
		go printFruit(fruit, i, &wg)
	}

	// Tunggu semua goroutine selesai
	wg.Wait()

	fmt.Println("Semua goroutine selesai dijalankan.")
}
