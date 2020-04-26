/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	//暴力法 笨办法... 时间O(n) 空间O(n)
/*  stack := []byte{}
	index := []int{}
	for k, v := range []byte(S) {
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if v == ')' {
			stack = stack[0: len(stack) - 1]
		}
		sl := len(stack) 
		if sl == 0 {	//发现栈长度为0 记录k的位置 说明是配对的地方
			index = append(index, k)
		}
	}
	s := ""
	next := 0
	for _, v := range index {
		tmp := S[next:v + 1]	//根据位置分割s
		next = v + 1
		s += tmp[1:len(tmp) - 1] //移除最外层的括号并重新拼接
	}
	return s */
	//好像可以不用栈
 /* index := []int{}
	pos := 0
	for k, v := range []byte(S) {
		if v == '(' {
			pos++
		} else {
			pos--
		}
		if pos == 0 {
			index = append(index, k)
		}
	}
	s := ""
	next := 0
	for _, v := range index {
		tmp := S[next:v + 1]	//根据位置分割s
		next = v + 1
		s += tmp[1:len(tmp) - 1] //移除最外层的括号并重新拼接
	}
	return s */
	//再优化一下
	pos := 0
	tmp, s := "", ""
	for _, v := range []byte(S) {
		tmp += string(v)
		if v == '(' {
			pos++
		} else {
			pos--
		}
		if pos == 0 {
			s += tmp[1:len(tmp)-1]
			tmp = ""
		}
	}
	return s
}
// @lc code=end

