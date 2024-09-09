package main

import (
	"fmt"
	"goproject/ch14/ex14.2/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()
	
	var myint publicpkg.MyInt = 10
	fmt.Println("myint: ", myint)
}