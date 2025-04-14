package session2

import "fmt"

func Sesi26() {
	str := "mânca"

	for i, s := range str {
		fmt.Printf("Index : %d, string : %s\n", i, string(s))
	}
}
