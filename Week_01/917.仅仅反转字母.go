/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	l := len(S)
	if l == 0 {
		return S
	}
	//双指针试试 时间复杂度O(n) 空间复杂度O(n)
	left, right := 0, l - 1
	s := []byte(S)
	for left < right {
		if !isLetter(s[left]) {
			left++
			continue
		}
		if !isLetter(s[right]) {
			right--
			continue
		}
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}
func isLetter (str byte) bool{
    if (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') {
        return true
    }
    return false
}
// @lc code=end

