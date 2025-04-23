package session5

import (
	"errors"
	"fmt"
)

func validPassword(password string) (string, error) {
	if len(password) < 4 {
		return "Valid password", nil
	}
	return "", errors.New("Password harus kurang dari 4 karakter")
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi panic:", r)
	}
}

func Sesi52() {

	defer catchErr() // recover() akan menangkap panic di sini
	var password string
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	result, err := validPassword(password)
	if err != nil {
		//fmt.Println("Error:", err)
		panic(err.Error())
	} else {
		fmt.Println(result)
	}
}
