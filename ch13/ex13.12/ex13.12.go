package main

import "fmt"

func main() {
	str1 := "AAA"
	str2 := "aaabtc"

	// UTF-8 코드의 정수값으로 비교를 하기 때문에 소문자 a가 97번이고 A가 65번이기에 a가 더 크다구 합니다.
	fmt.Printf("%s > %s => %v\n", str1, str2, str1 > str2)
}