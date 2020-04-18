package main

import (
	"fmt"
)

func main() {
	input := 1230321
	output := isPalindrome(input)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	arr := []int{}

	for {
		if x == 0 {
			break
		}

		c := x % 10

		arr = append(arr, c)

		x = x / 10
	}

	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-i-1] {
			return false
		}
	}

	return true
}

// @lc code=end




