/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	//暴力法来一波
	//ln := len(nums)
	/*
	if ln == 0 {
		return -1
	}
	hmap := make(map[int]int, ln)
	for _, v := range nums {
		hmap[v]++
		if hmap[v] > (ln / 2) {
			return v
		}
	}
	return -1
	*/
	//排序的来一波
	/*
	sort.Ints(nums)
	return nums[ln/2]
	*/
	// 投票算法来一波
	count := 0
	var candidate int 
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		}else{
			count--
		}
	}
	return candidate
}
// @lc code=end

