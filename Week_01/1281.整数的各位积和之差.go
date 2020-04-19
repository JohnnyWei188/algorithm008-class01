/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	/* 
	 暴力法来一波 时间复杂度 0(n)
	if n <= 0 || n > 10e5{
		return 0
	}
	strN := strconv.Itoa(n)
	sum, product := 0, 1
	for _, v := range strN {
		iv, _ := strconv.Atoi(string(v))
		sum +=iv
		product *=iv
	}
	return product - sum
	*/
	//有没有更佳的方法 
	sum, mul := 0, 1
	for n > 0 {
		digit := n % 10  // 取个位数的值
		n /= 10   
		sum += digit
		mul *= digit
	}
	return mul - sum
}
// @lc code=end

