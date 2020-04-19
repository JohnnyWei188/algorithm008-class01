/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	//双指针法
	l := len(nums)
	if l == 0 {
		return 0
	}
	i := 0
	for j := 1; j < l; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}
// @lc code=end

