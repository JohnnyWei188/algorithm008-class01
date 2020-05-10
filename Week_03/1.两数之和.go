/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int, 0)
	ret := make([]int, 0)
	for k, v := range nums {
		complement := target - v	//
		vv, ok := hmap[complement]
		if ok {
			return []int{vv, k}
		}
		hmap[v] = k
	}
	return ret
}
// @lc code=end

