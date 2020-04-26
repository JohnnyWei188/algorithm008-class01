/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	//迭代
	ret := make([]int, 0)
	if root == nil {
		return ret 
	}
	var f func(*TreeNode)
	f = func(n *TreeNode){
		if n == nil {
			return
		}
		ret = append(ret, n.Val)
		f(n.Left)
		f(n.Right)
	}
	f(root)
	return ret
}
// @lc code=end

