package main

import "fmt"

func main() {
	var a int = 500
	var p *int
	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("a의 주소값: %p\n", &a)
	fmt.Printf("p의 주소값: %p\n", &p)
	fmt.Printf("p가 가리키는 값: %d\n", *p)
}