package session5

import (
	"fmt"
	"time"
)

func Sesi5() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine pertama mengirim data ke channel c1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Data dari c1"
	}()

	// Goroutine kedua mengirim data ke channel c2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Data dari c2"
	}()

	// Menerima data dari kedua channel menggunakan select
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Menerima:", msg1)
		case msg2 := <-c2:
			fmt.Println("Menerima:", msg2)
		}
	}
}
