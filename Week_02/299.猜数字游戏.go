/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	//暴力法来一波
	l := len(secret)
	if l == 0 {
		return ""
	}
	hmap := map[string]int{}
	bulls, cows := 0, 0
	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {	//相等的话, 就是bulls的个数
			bulls++
			continue
		}
		hmap[string(secret[i])]++ //这里一加
		hmap[string(guess[i])]--  //这里一减 就是统计猜中了的字符
		fmt.Println(hmap)
	}
	for _, v := range hmap {
		if v > 0 {	// v 大于 0 表示的是secret中没有猜中的数字的个数
			cows += v
		}
	}
	fmt.Println(hmap)
	fmt.Println(cows)
	return fmt.Sprintf("%dA%dB", bulls, l - cows - bulls)
}

//计算bucket中正值的个数，其实就是secret中没有匹配到的数字的个数，
//length-bulls就是guess中除去猜对数字位置的个数，
//再减去没有匹配到的数字的个数，就是cows
// @lc code=end

