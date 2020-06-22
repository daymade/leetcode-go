package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

	children := []*Node{
		&Node{
			Val: 2,
			Children: nil,
		}
	}

	root := &Node{
		Val: 1,
		Children: children
	}

	depth := maxDepth(root)
	print(depth)
}

/*
 * @lc app=leetcode id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	return dfs(root)
}

func dfs(node *Node) int {
	if node == nil {
		return 0
	}

	md := 1
	for _, child := range node.Children {
		d := dfs(child) + 1
		md = int(math.Max(float64(md) , float64(d)))
	}

	return md;
}

// @lc code=end

