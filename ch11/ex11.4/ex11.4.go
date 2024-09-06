package main

import "fmt"

type User struct {
	Name	string
	ID		string
	Age		int
	Level	int
}

type VIPUser struct {
	User
	Price	int
	Level	int
}

func main() {
	user := User{ "김갑환", "Hwan", 25, 10 }
	vip := VIPUser{
		User{ "화랑", "hwarang", 30, 11 },
		250,
		3, // 여러 줄로 초기화할 때는 가장 마지막 값 뒤에 꼭 쉼표를 달기
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.User.Level,
	)
}