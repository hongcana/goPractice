package main

import "fmt"

func main() {
	str1 := "Never"
	str2 := "Sell"
	str3 := "Your"
	str4 := "ETH"
	
	str5 := str1 + " " + str2 + " " + str3 + " " + str4
	fmt.Println(str5)

	str3 += " " + str4
	fmt.Println(str3)
}