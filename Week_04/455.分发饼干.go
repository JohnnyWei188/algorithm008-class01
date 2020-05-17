/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	lg := len(g)
	ls := len(s)
	if lg == 0 || ls == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	n := 0 
	i := 0
	j := 0
	for i < ls {
		if j >= lg {
			return n
		}
		if s[i] >= g[j] {
			n++
			j++
		}
		i++
	}

	return n
}
// @lc code=end

