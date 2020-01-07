package main

import "fmt"

func main() {
	ans := twoSum([]int{2, 7, 11, 15}, 26)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	/*
		Given nums = [2, 7, 11, 15], target = 9,
		Because nums[0] + nums[1] = 2 + 7 = 9,
		return [0, 1].
	*/
	for width := 1; width < len(nums)-1; width++ {
		for idx, _ := range nums[0 : len(nums)-width] {
			if target == nums[idx]+nums[idx+width] {
				return []int{idx, idx + width}
			}
		}
	}
	return nil
}
