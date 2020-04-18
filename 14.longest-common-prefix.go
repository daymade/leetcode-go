package main

import (
	"fmt"
	"math"
)

func main() {
	input := []string{"flower", "flow", "flight"}
	output := longestCommonPrefix(input)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lcp := ""

	l := len(strs[0])

	for _, str := range strs {
		l = int(math.Min(float64(l), float64(len(str))))
	}

	for i := 0; i < l; i++ {
		c := strs[0][i]
		for _, str := range strs {
			if c != str[i] {
				return lcp
			}
		}
		lcp = lcp + string(c)
	}

	return lcp
}

// @lc code=end
