package main

//still challenging...
import "fmt"

func main() {
	a := reverse(123)
	fmt.Println(a)
}

func reverse(x int) (ans string) {
	//var quotient int
	var nums []int
	if x < 0 {
		x = -x
	}
	for x > 0 {
		remainder := x % 10
		nums = append(nums, remainder)
		x = x / 10
	}
	//	for i, num := range {
	//		ans :=
	ans = fmt.Sprintf("%d", nums)
	return ans
}

func pow(n int) (m int) {
	if n >= 1 {
		m *= 10
		n--
		pow(n)
	}
	return m
}
