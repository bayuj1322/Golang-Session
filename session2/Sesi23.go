package session2

import "fmt"

func Sesi23() {
	cars := []string{"BMW", "Honda", "Toyota", "Ford"}
	fmt.Println("cars =>", cars)

	// newCars copy dari cars[0:2] tapi make append → bikin array baru
	newCars := append([]string{}, cars[0:2]...)
	fmt.Println("newCars =>", newCars)

	// ubah data di cars
	cars[0] = "Tesla"

	fmt.Println("\nSetelah cars[0] diubah jadi Tesla:")
	fmt.Println("cars =>", cars)
	fmt.Println("newCars =>", newCars)
}
