package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Score	float64
	Age		int32
	Age2	int32
}

func main() {
	user := User {32.8, 64, 1}
	fmt.Println(unsafe.Sizeof(user))
}