package session3

import (
	"fmt"
	"strings"
)

func profile(name string, favFoods ...string) {
	mergefavfood := strings.Join(favFoods, ",")

	fmt.Println("Nama:", name)
	fmt.Println("Makan:", mergefavfood)

}

func Sesi3() {
	// Pemanggilan fungsi dengan 1 parameter biasa dan beberapa variadic
	profile("Airell", "Nasi Goreng", "Sate", "Bakso", "Mie Ayam")
}
