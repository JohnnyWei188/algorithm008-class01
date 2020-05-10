/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
	//递归
	/*
	if root == nil {
		return 0
	}
	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)
	return max(leftDep, rightDep) + 1
	*/
	//BFS
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := 0
	for len(queue) > 0 {
		leveSize := len(queue)
		for i := 0; i < leveSize; i++ {
			s := queue[0]
			queue = queue[1:]
			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		ret++
	}
	return ret
}
func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

