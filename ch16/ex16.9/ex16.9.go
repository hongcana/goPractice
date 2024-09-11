package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1) // 첫번째 인자로 복사한 결과를 저장하는 슬라이스 변수, 두 번째는 source, return은 복사된 요쇼의 개수

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}