package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "궁내 체고의 싱카볼 투순데.."
	var slice []byte = []byte(str)
	var runes []rune = []rune(str)
	
	fmt.Printf("%s str:\t%p\n", str, unsafe.StringData(str))
	fmt.Printf("%s slice:\t%p\n", slice, unsafe.SliceData(slice))
	fmt.Printf("%c runes:\t%p\n", runes, unsafe.SliceData(runes))
}