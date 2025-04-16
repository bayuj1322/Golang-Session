package session3

import "fmt"

// Struct biasa
type Person struct {
	Name string
	Age  int
}

func Sesi33() {
	// Anonymous struct tanpa pengisian nilai field
	employee1 := struct {
		ID     int
		Role   string
		Person Person
	}{} // <- inisialisasi harus tetap pakai kurung kurawal meskipun kosong

	employee1.ID = 1
	employee1.Role = "Manager"
	employee1.Person.Name = "Julian"
	employee1.Person.Age = 17

	fmt.Printf("Employee 1: %+v\n", employee1)

	// Anonymous struct dengan pengisian nilai field
	employee2 := struct {
		ID     int
		Role   string
		Person Person
	}{
		ID:   101,
		Role: "Developer",
		Person: Person{
			Name: "Alice",
			Age:  28,
		},
	}

	fmt.Printf("Employee 2: %+v\n", employee2)
}
