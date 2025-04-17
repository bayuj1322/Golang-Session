package session4

import (
	"fmt"
	"reflect"
)

// Struct student dengan field dan method public
type student struct {
	Name string
}

// Method SetName untuk mengubah nama student
func (s *student) SetName(newName string) {
	s.Name = newName
}

func Sesi42() {
	// Buat instance student
	s1 := &student{Name: "john wick"}
	fmt.Println("Sebelum:", s1.Name)

	// Ambil reflect.Value dari objek s1
	v := reflect.ValueOf(s1)

	// Cari method bernama "SetName"
	method := v.MethodByName("SetName")
	if method.IsValid() {
		// Buat parameter yang akan dikirim (harus dalam bentuk []reflect.Value)
		params := []reflect.Value{reflect.ValueOf("ethan hunt")}

		// Panggil method SetName dengan parameter
		method.Call(params)

		fmt.Println("Setelah:", s1.Name)
	} else {
		fmt.Println("Method SetName tidak ditemukan.")
	}
}
