package main

import (
	"fmt"
	"math"
)

func main() {
	input := -1230
	output := reverse(input)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	r := 0

	for {
		if x == 0 {
			if r > math.MaxInt32 || r < math.MinInt32 {
				return 0
			}

			return r
		}

		r = r * 10

		c := x % 10
		x = x / 10

		r = r + c
	}
}

// @lc code=end
