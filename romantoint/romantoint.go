package main

import "fmt"

func main() {
	a := romanToInt("MCMXCIV") //1994
	fmt.Printf("Answer is %d.\n", a)
	a = romanToInt("III")
	fmt.Printf("Answer is %d.\n", a)
	a = romanToInt("IV")
	fmt.Printf("Answer is %d.\n", a)
	a = romanToInt("CMLII")
	fmt.Printf("Answer is %d.\n", a) //952
}

func romanToInt(s string) int {
	var ans int
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	specialM := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	isPreviousSpecial := false
	isLastCalculated := false
	for i := 0; i < len(s)-1; i++ {
		isLastCalculated = false
		if i > 0 {
			_, isPreviousSpecial = specialM[s[i-1:i+1]]
		}
		if isPreviousSpecial != true {
			_, isSpecial := specialM[s[i:i+2]]
			if isSpecial == true {
				ans += specialM[s[i:i+2]]
				isLastCalculated = true
			} else {
				ans += m[s[i:i+1]]
			}
		}
	}
	if isLastCalculated != true {
		ans += m[s[len(s)-1:len(s)]]
	}
	return ans
}
