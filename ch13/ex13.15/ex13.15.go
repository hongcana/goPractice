package main

import "fmt"

func main() {
	str1 := "궁내 체고의 싱카볼 투순데"
	fmt.Printf("%v, %s\n", &str1, str1)
	str1 = "오다가 직각으로 하나 떨어져주면 좋은데요."
	fmt.Printf("%v, %s\n", &str1, str1)
}  