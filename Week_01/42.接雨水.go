/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	//暴力法不会， 看了答案抄的
	sum := 0
	stack := make([]int, 0)
	current := 0
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			h := height[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			min := min(height[stack[len(stack)-1]], height[current])
			sum += distance * (min - h)
		}
		stack = append(stack, current)
		current++
	}
	return sum
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

