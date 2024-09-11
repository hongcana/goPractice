package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age int
}

type Students []Student

func (s Students) Len() int { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"양의지", 37},
		{"김재환", 35},
		{"양석환", 33},
		{"이유찬", 31},
		{"김대한", 25},
		{"김택연", 21},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}