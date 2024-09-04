package main

import "fmt"

type Book struct {
	Title string
	Author string
	Year int
	Price float64
}

func (b Book) DiscountedPrice(discount float64) float64 {
	return b.Price - (b.Price * discount)
}

func (b Book) DisplayInfo() {
	fmt.Printf("제목: [%s]\n",b.Title)
	fmt.Printf("저자: [%s]\n",b.Author)
	fmt.Printf("출판 연도: [%d]\n",b.Year)
	fmt.Printf("가격: [%.1f]\n",b.Price)
	fmt.Printf("할인 가격: [%.1f]\n",b.DiscountedPrice(0.1))
}

func main() {
	book := Book{
		Title: "Go 프로그래밍",
		Author: "홍길동",
		Year: 2023,
		Price: 30000.0,
	}
	book.DisplayInfo()
}