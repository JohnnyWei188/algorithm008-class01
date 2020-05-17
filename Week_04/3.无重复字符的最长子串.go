/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	max := 0
	j := 0
	hmap := map[byte]int{}
	for i := 0; i < l; i++ {
		v, ok := hmap[s[i]]
		if ok {
			//j表示的是左侧的指针，maxInt表示的是确保
			//指针往右边移动
			j = maxInt(j, v + 1)
		}
		hmap[s[i]] = i
		//max始终记录的是最长的子串
		max = maxInt(max, i - j + 1)
	}
	return max
}
func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

