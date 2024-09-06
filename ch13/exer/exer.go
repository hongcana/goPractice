package main

import (
	"fmt"
	"unsafe"
)

type User struct{
	FirstName		string
	LastName		string
	Age				int
}

func main() {
	user := User{ 
		"string",
		"string",
		28, 
	}
	fmt.Println(unsafe.Sizeof(user))
}