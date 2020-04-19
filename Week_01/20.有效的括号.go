/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	//暴力法来一波
	/*
		l := len(s)
		if l == 0 {
			return true
		}
		target := []string{"()", "[]", "{}"}
		for i := 0; i < l - 1; i++ {
			if s == "" {
				break
			}
			for _, v := range target {
				s = strings.ReplaceAll(s, v, "")
			}
		}
		if len(s) == 0 {jij
			return true
		}
		return false
	*/
	//栈来一波
	l := len(s)
	if l == 0 {
		return true
	}
	braMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := []string{}
	for _, v := range s {
		sv := string(v)
		for _, vv := range braMap { //发现有左括号, 入栈
			if sv == vv {
				stack = append(stack, sv)
			}
		}
		slen := len(stack)
		if bv, ok := braMap[sv]; ok { //有右括号
			if slen == 0 { //栈已经为空，还需要出栈，说明肯定是不匹配的
				return false
			}
			if bv != stack[slen-1] { //与需要出栈的元素不匹配
				return false
			}
			stack = stack[0 : slen -1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end
