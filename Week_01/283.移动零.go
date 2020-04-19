/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	l := len(nums)
	if l == 0 {
		return 
	}
	posIndex := 0
	for k, v := range nums {
		if v != 0 {
			nums[posIndex], nums[k] = nums[k], nums[posIndex]
			posIndex++
		}
	}
}
// @lc code=end

