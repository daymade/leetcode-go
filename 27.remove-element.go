package main

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	len := removeElement(nums, 2)
	for i := 0; i < len; i++ {
		print(nums[i])
	}
}

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	p1 := 0
	p2 := 0
	for p2 < len(nums) {
		if nums[p2] != val {
			nums[p1] = nums[p2]
			p1++
		}
		p2++
	}

	return p1
}

// @lc code=end
