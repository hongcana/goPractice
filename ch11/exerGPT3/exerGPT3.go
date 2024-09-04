package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // 함수 구현
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target int
	_, err := fmt.Scan(&target)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(twoSum(nums, target))
}