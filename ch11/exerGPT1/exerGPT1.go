package main

import "fmt"

type Student struct {
	Name string
	Age int
	Grade float64
}

func (s Student) isPassed() bool {
	if s.Grade >= 60.0 {
		return true
	} else {
		return false
	}
}

func (s Student) DisplayInfo() {
	fmt.Printf("이름: [%s]\n", s.Name)
	fmt.Printf("나이: [%d]\n", s.Age)
	fmt.Printf("성적: [%.1f]\n", s.Grade)
	fmt.Printf("합격 여부: [%t]\n", s.isPassed())
}

func main() { 
	foo := Student{
		Name: "홍길동",
		Age: 20,
		Grade: 75.5,
	}
	foo.DisplayInfo()
}
