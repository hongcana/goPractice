package main

import "fmt"

const Y int = 3

func main() {
	x := 5
	// a := [x]int{1,2,3,4,5} // 에러! 변수를 배열의 길이 초기화에 사용할 수 없음.
	a := [Y]int{1,2,3} // 상수는 사용 가능함.

	var b [10]int

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
}