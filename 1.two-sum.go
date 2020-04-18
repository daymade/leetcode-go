package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	output := twoSum(input, 4)
	fmt.Println(output)
}

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for index, item := range nums {
		item2 := target - item

		if index2, ok := m[item2]; ok {
			return []int{index2, index}
		}

		m[item] = index
	}
	panic("no answer")
}

// @lc code=end
