/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	//暴力法
	l := len(prices)
	if l == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < l - 1; i++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
// @lc code=end

