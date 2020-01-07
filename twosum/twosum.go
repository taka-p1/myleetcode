package main

import "fmt"

func main() {
	ans := twoSum([]int{2, 7, 11, 15}, 26)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	for width := 1; width <= len(nums)-1; width++ {
		for idx, _ := range nums[0 : len(nums)-width] {
			if target == nums[idx]+nums[idx+width] {
				return []int{idx, idx + width}
			}
		}
	}
	return nil
}
