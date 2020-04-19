/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	//暴力来一波  O(nlogn) 时间复杂度
	/*
	for k, v := range A {
		A[k] *= v
	}
	sort.Ints(A) //这里有个排序
	return A	
	//双指针
	l := len(A)
	if l == 0 {
		return A
	}
	*/
	//双指针法
	l := len(A)
	r := make([]int, l)
	left, right := 0, l - 1
	for p := l -1; p >= 0; p-- {
		if abs(A[left]) > abs(A[right]) {
			r[p] = A[left] * A[left]
			left++
		} else {
			r[p] = A[right] * A[right]
			right--
		}
	}
	return r
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
// @lc code=end

