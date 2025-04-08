package session2

import "fmt"

func Sesi2() {
	// IF ELSE Statement
	nilai := 85

	fmt.Println("== Contoh IF ELSE ==")
	if nilai >= 90 {
		fmt.Println("Grade: A")
	} else if nilai >= 80 {
		fmt.Println("Grade: B")
	} else if nilai >= 70 {
		fmt.Println("Grade: C")
	} else if nilai >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}

	//switch case

	Nilai := 8
	fmt.Println("\n== Contoh SWITCH CASE ==")
	switch true {
	case Nilai <= 5:
		fmt.Println("Keterangan: Kurang")
	case Nilai == 6 || Nilai == 7:
		fmt.Println("Keterangan: Cukup")
	case Nilai == 8 || Nilai == 9:
		fmt.Println("Keterangan: Bagus")
	case Nilai == 10:
		fmt.Println("Keterangan: Sempurna")
	default:
		fmt.Println("Skor tidak valid")
	}
}
