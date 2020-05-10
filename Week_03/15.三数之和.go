/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	//双指针法
	ret := make([][]int, 0)
	ln := len(nums)
	if ln == 0 {
		return ret
	}
	sort.Ints(nums)
	left, right := 0, ln - 1
	for i := 0; i < ln; i++ {
		if nums[i] > 0 { //最左数都大于0 不可能存在 a + b + c = 0的情况
			return ret
		}
		if i > 0 && nums[i] == nums[i - 1] { //去除重复的情况
			continue
		}
		left = i + 1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				ret = append(ret, tmp)
				for left < right && nums[left] == nums[left+1] {
					left++ //移动left指针 去重
				}
				for left < right && nums[right] == nums[right-1]{
					right--
				}
				left++
				right--
			}else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			}else {
				left++
			}
		}
	}
	return ret
}
// @lc code=end

