/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	//递归
	/*
	ret := make([]int, 0)
	var f func(*Node)
	f = func(n *Node){
		if n == nil {
			return 
		}
		ret = append(ret, n.Val)
		for _, v := range n.Children {
			f(v)
		}
	}
	f(root)
	return ret
	*/
	//迭代
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := []*Node{root}	//通过数组来模拟一个栈
	for len(stack) > 0 {
		node := stack[len(stack) - 1] 	//弹出一个
		stack = stack[:len(stack) - 1]
		for i := len(node.Children) -1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		ret = append(ret, node.Val)
	}
	return ret
}

// @lc code=end

