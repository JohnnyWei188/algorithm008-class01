/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := len(nums) 
	if l == 0 {
		return nums
	}
	//暴力的来一波
	/*
	for i := 0; i < l - 1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
	*/
	//一遍哈希表
	hp := map[int]int{}
	for i := 0; i < l; i++ {
		index, ok := hp[target - nums[i]]
		if ok {
			return []int{index, i}
		}
		hp[nums[i]] = i
	}
	return []int{}
}
// @lc code=end

