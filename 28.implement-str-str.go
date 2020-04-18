package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "pi"
	i := strStr(haystack, needle)
	fmt.Println(i)
}

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			eq := true
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					eq = false
				}
			}
			if eq {
				return i
			}
		}
	}

	return -1
}

// @lc code=end
