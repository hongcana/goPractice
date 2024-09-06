package main

import "fmt"

func main() {
	str1 := "Hello\t World!\n"
	
	str2 := `Go is "awesome!"\t\n`
	fmt.Println(str1)
	fmt.Println(str2)
}