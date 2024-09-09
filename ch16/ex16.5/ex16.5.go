package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5) // [0, 0, 0]

	slice2 := append(slice1, 4, 5) // [0, 0, 0, 4, 5]

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // slice2까지 바뀝니다. 슬라이스의 내부 배열은 배열을 가리키는 포인터의 값을 복사하게 됩니다.
	
	fmt.Println("After change second ele.")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1)) // [0, 100, 0]
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2)) // [0, 100, 0, 4, 5]

	slice1 = append(slice1, 500) // 역시나 slice2까지 바뀝니다.
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1)) // [0, 100, 0, 500]
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2)) // [0, 100, 0, 500, 5]
}