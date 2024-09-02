package main

import "fmt"

func main() {
	var a float32 = 12345.523
	var b float32 = 34567.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}