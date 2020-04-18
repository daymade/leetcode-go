package main

import (
	"fmt"
)

func main() {
	input := "MCMXCIV"
	output := romanToInt(input)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0

	prev, _ := m[s[0]]

	for i := 1; i < len(s); i++ {
		num, _ := m[s[i]]

		if prev < num {
			sum = sum - prev
		} else {
			sum = sum + prev
		}

		prev = num
	}

	return sum + prev
}
// @lc code=end

