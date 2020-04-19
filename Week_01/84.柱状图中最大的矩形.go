/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return 1 * heights[0]
	}
	maxArea, area := 0, 0
	for i := 0; i < l - 1; i++ {
		for j := i + 1; j < l; j++ {
			area = min(heights[i], heights[j]) * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
// @lc code=end

