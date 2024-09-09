package publicpkg

import "fmt"

const (
	PI = 3.141592 // 공개되는 상수
	pi = 3.1416   // 공개되지 않음
)

var ScreenSize int = 1080
var screenHeight int

func PublicFunc() {
	const MyConst = 100 // 공개되는 함수 내부라도 함수 내부에서 상수를 선언하면 패키지 외부로 공개되지 않음
	fmt.Println("This is a public Function", MyConst)
}

func privateFunc() {
	fmt.Println("공개되지 않습니다.(소문자 시작)")
}

type MyInt int
type myString string

type MyStruct struct {	// 공개
	Age	int				// 공개
	name	string		// 비공개
}

func (m MyStruct) PublicMethod() {
	fmt.Println("this is public method")
}

func (m MyStruct) privateMethod() {
	fmt.Println("this isn't public method", m.name)
}