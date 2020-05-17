/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
 var ret [][]int
 var visited []*TreeNode
func levelOrder(root *TreeNode) [][]int {
	//DFS 
	ret = [][]int{}
	visited = []*TreeNode{}
	level := 0
	dfs(root, level) //深度优先 每次传入level
	return ret
}
func dfs(root *TreeNode, level int){
	if inArray(root, visited) {
		return  
	}
	if level == len(ret) {
		ret = append(ret, []int{})
	}
	ret[level] = append(ret[level], root.Val)
	if root.Left != nil {
		dfs(root.Left, level + 1)
	}
	if root.Right != nil  {
		dfs(root.Right, level + 1)
	}
}

func inArray(node *TreeNode, arr []*TreeNode) bool {
	for _, v := range arr {
		if node == v {
			return true
		}
	}
	return false
}

// @lc code=end

