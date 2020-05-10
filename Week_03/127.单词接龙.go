/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//暴力法试试
	if beginWord == endWord {
		return 1
	}
	refMap := make(map[string]bool, 0)
	for _, v := range wordList {
		refMap[v] = true
	}
	
}
// @lc code=end

