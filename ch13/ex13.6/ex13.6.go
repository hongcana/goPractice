package main

import "fmt"

func main() {
	// string type과 []rune type은 모두 문자들의 집합을 나타내기 때문에 상호 타입 변환이 가능합니다.
	// UTF8에서 한글은 문자당 3바이트, 영문은 1바이트. so len of str = 12bytes
	// string type을 []rune type으로 변환하면 각 글자들로 이뤄진 배열로 변환됨.
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(runes) = %d\n", len(runes))
}