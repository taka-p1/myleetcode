package main

import "fmt"

func main() {
	fmt.Println(isPalindorome(12321))
}

func isPalindorome(x int) bool {
	origin := x
	if x < 0 {
		return false
	}
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev == origin
}
