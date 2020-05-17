/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	if _, ok := dict[endWord]; !ok { //加入到字典中
		return 0
	}
	queue := []string{beginWord}
	l := len(beginWord)
	steps := 0
	for len(queue) > 0 {
		steps++
		size := len(queue)
		for i := size; i > 0; i-- {
			s := queue[0]
			queue = queue[1:]
			chs := []rune(s)
			for i := 0; i < l; i++ {
				ch := chs[i]
				for c := 'a'; c <= 'z'; c++ {
					if c == ch {
						continue
					}
					chs[i] = c 
					t := string(chs)
					if t == endWord {
						return steps + 1
					}
					if _, ok := dict[t]; !ok {
						continue
					}
					delete(dict, t)
					queue = append(queue, t)
				}
				chs[i] = ch	//将单词的第i位复位 
			}
			fmt.Println(queue)
		}
	}
	return 0
}
// @lc code=end

