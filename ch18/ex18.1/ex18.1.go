package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name	string
	Age		int
}

func (s Student) String() string {
	// 얘는 반환하는 애임
	return fmt.Sprintf("안뇽 나는 %d살 %s라고 해.", s.Age, s.Name)
}

func main() {
	student := Student{ "짱구", 23 }
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())
}