/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	queue := []string{start}
	count := 0
	now := 0
	top := 1
	for len(queue) > 0 {
		count++
		for _, ss := range queue {
			for {
				if top < 0 {
					break
				}
				for i := 0; i < len(bank); i++ {
					s := bank[i]
					if s == "" {
						continue
					}
					if !compare(ss, s) {
						continue
					}else if s == end {
						return count
					}
					bank[i] = ""
					queue = append(queue, s)
					now++
				}
			}
		}
		top--
	}
	return -1
}
func compare(start string, s string) bool {
	count := 0
	for i := 0; i < len(start); i++ {
		if start[i:i+1] != s[i: i+1] {
			count++
		}
	}
	return count == 1
}
// @lc code=end

