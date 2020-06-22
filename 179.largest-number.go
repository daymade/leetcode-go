package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	num := largestNumber(nums)
	print(num)
}

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) string {

	sort.SliceStable(nums, func(i, j int) bool {
		ns1 := strconv.Itoa(nums[i])
		ns2 := strconv.Itoa(nums[j])

		return ns1+ns2 > ns2+ns1
	})

	s := ""

	for _, v := range nums {
		s = s + strconv.Itoa(v)
	}

	if strings.Count(s, "0") == len(s) {
		return "0"
	}

	return s
}

// @lc code=end
