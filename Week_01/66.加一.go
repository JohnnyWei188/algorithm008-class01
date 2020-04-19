/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return digits
	}
	for i := l - 1; i >= 0; i-- {
		digits[i]++	//每个位都加1
		digits[i] = digits[i] % 10 //判定是否进位
		if digits[i] != 0 {	//如果数组没有把位数都循环完，说明不是99,999这样的数 
			return digits
		}
	}
	//否则 如果有进位 如99 这种 
	digits = make([]int, l + 1)
	digits[0] = 1 
	return digits
}
// @lc code=end

