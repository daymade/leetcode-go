package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len := removeDuplicates(nums)
	for i := 0; i < len; i++ {
		print(nums[i])
	}
}

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	p := nums[0]
	l := 0
	for i := 1; i < len(nums); i++ {
		if p != nums[i] {
			l++
		}
		nums[l] = nums[i]
		p = nums[i]
	}

	return l + 1
}
// @lc code=end

