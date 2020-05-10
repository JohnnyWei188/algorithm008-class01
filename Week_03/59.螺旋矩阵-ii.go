/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	//暴力法试试
	ret := make([][]int, n)
	if n == 0 {
		return ret
	}
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	num, max := 1, n * n
	left, right, top, buttom := 0, n - 1, 0, n -1 
	for num <= max {
		for i := left; i <= right; i++ {	//从矩阵top端开始写 从左到右
			ret[top][i] = num
			num++
		}
		top++	//矩阵的top边界 下移1个位置
		for i := top; i <= buttom; i++ {	//矩阵最右边 从上写到下
			ret[i][right] = num
			num++
		}
		right-- //矩阵的右边界 向左移动1个位置
		for i := right; i >= left; i-- { //从右往左边写
			ret[buttom][i] = num
			num++
		}
		buttom-- //矩阵的底边界 上移1个位置
		for i := buttom; i >= top; i-- {
			ret[i][left] = num
			num++
		}
		left++
	}
	return ret
}
// @lc code=end

