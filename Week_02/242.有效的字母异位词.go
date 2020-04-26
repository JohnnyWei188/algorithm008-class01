/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	//暴力法 
	/* if len(s) != len(t) {
		return false
	}
 	sb := []byte(s)
	st := []byte(t)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.Slice(st, func(i, j int) bool {
		return st[i] < st[j]
	})
	s = string(sb)
	t = string(st)
	if s == t {
		return true
	}
	return false */
	//使用hash tab的方式
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}
	hmap := map[byte]int{}
	for i := 0; i < ls; i++ {
		hmap[s[i]]++
	}
	for j := 0; j < lt; j++ {
		hmap[t[j]]--
		if hmap[t[j]] < 0 {
			return false
		}
	}
	return true
}
// @lc code=end

