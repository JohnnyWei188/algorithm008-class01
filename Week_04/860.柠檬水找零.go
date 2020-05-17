/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0 
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else { //20 元的
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false 
			}
		}
	}
	return true
}
// @lc code=end

