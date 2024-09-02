package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 123.1, 3523.3, 102.8}
	
	for i, v := range t { // 만약 선언하고 사용하지 않는 변수가 있으면 컴파일시 에러가 발생함.
		fmt.Println(i, v)
	}
}