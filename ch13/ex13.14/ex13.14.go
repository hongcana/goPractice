package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len int
}

func main() {
	// 과연 고언어에서 스트링 타입의 복사는 value 복사일까요, Reference 복사일까요?
	str1 := "궁내 체고의 싱카볼 투순데"
	var str2 string = str1

	StringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	StringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(StringHeader1)
	fmt.Println(StringHeader2)
}