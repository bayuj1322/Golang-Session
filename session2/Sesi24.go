package session2

import "fmt"

type Pegawai struct {
	Nama string
	Age  int
}

func Sesi24() {
	// Slice berisi daftar pegawai
	pegawais := []Pegawai{
		{Nama: "Ayu", Age: 22},
		{Nama: "Budi", Age: 23},
		{Nama: "Citra", Age: 24},
	}
	for i, value := range pegawais {
		fmt.Printf("index: %d,Nama: %s, Age: %d\n", i, value.Nama, value.Age)
	}
}
