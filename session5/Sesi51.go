package session5

import (
	"fmt"
	"os"
)

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func callDeferFunc() {
	defer deferFunc()
	fmt.Println("callDeferFunc is executing")
}

func Sesi51() {
	callDeferFunc()
	os.Exit(1)
	fmt.Println("Hai everyone!!")
}
