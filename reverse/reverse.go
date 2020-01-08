package main

import "fmt"

func main() {
	a := reverse(120)
	fmt.Println(a)
}

func reverse(x int) (ans int) {
	var nums []int
	tmp := x
	if x < 0 {
		x = -x
	}
	for x > 0 {
		remainder := x % 10
		nums = append(nums, remainder)
		x = x / 10
	}
	i := 0
	for n := len(nums) - 1; n >= 0; n-- {
		ans += nums[i] * pow(n)
		i++
	}
	if tmp < 0 {
		ans = -ans
	}
	return ans
}

func pow(n int) int {
	if n == 0 {
		return 1
	} else {
		return 10 * pow(n-1)
	}
}
