/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left)
	r := dfs(node.Right)

	return int(math.Max(float64(l), float64(r))) + 1
}
// @lc code=end

