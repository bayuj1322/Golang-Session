package session4

import (
	"fmt"
	"runtime"
	"time"
)

func firstProcess(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("firstProcess:", i)
	}
}

func secondProcess(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("secondProcess:", i)
	}
}

func Sesi43() {
	// Jalankan firstProcess sebagai goroutine
	go firstProcess(5)

	// Jalankan secondProcess secara biasa (tidak sebagai goroutine)
	secondProcess(5)

	// Tampilkan jumlah Goroutine yang sedang berjalan
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	// Tahan eksekusi main selama 2 detik agar goroutine sempat jalan
	time.Sleep(2 * time.Second)

}
