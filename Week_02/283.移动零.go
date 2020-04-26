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
	i, j := 0, 0
	for i = 0; i < l; i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
// @lc code=end

