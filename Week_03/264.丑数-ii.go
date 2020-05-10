/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		res[i] = min(res[p2] * 2, res[p3] * 3, res[p5] * 5)
		if res[i] == res[p2] * 2 {
			p2++
		}
		if res[i] == res[p3] * 3 {
			p3++
		}
		if res[i] == res[p5] * 5 {
			p5++
		}
	}
	return res[n - 1]
}

func min(a, b, c int) int {
	m := a 
	if a > b {
		m = b
	}
	if m > c {
		return c
	}
	return m
} 

// @lc code=end

