/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	/*
	暴力破解法  时间复杂度O(n^2) 空间复杂度O(1)
	l := len(nums)
	if l == 0 {
		return nums
	}
	for i := 0; i < l; i++ {
		v1 := nums[i]
		for j := i + 1; j < l; j++ {
			v2 := nums[j]
			if v1 + v2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
	*/
	//hash table 方式
	l := len(nums)
	if l == 0 {
		return nums
	}
	htab := make(map[int]int)
	for k, v := range nums {
		complement := target - v //计算目标对应的结果
		if idx, ok := htab[complement]; ok { //判定是否在htab中
			return []int{idx, k}
		}
		htab[v] = k
	}
	return []int{}
	
}
// @lc code=end

