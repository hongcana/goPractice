package main

import "fmt"

func F(n int) int {
	if (n == 1 || n == 0) {
		return n
	}
	return F(n-2) + F(n-1)
}

func main() {
	var a int

	fmt.Scan(&a)
	fmt.Println(F(a))
}