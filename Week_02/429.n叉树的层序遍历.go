/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := []*Node{root} //通过数组模拟一个队列  广度优先
	var level int
	for 0 < len(queue) {
		counter := len(queue)
		ret = append(ret, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				ret[level] = append(ret[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)	//入队
				}
			}
		}
		queue = queue[counter:] //出队
		level++
	}
	return ret 
}
// @lc code=end

