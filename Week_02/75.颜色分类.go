/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	//暴力法
	/*
	hmap := map[int]int{}
	index := []int{}
	for _, v := range nums {
		hmap[v]++
	}
	for k := range hmap {	//获取nums具体有哪些值，不再固定为0, 1, 2
		index = append(index, k)
	}
	sort.Ints(index)	//排序
	j := 0
	for k, v := range index {
		tmp := hmap[v]
		for tmp > 0 {
			nums[j] = index[k]
			j++
			tmp--
		}
	}
	*/
	//双指针 [2,0,2,1,1,0]
	left, curr, right := 0, 0, len(nums) - 1
	for curr <= right {
		if nums[curr] == 0 {
			nums[left], nums[curr] = nums[curr], nums[left]
			left++
			curr++
		}else if nums[curr] == 2 {
			nums[right], nums[curr] = nums[curr], nums[right]
			right--
		}else{
			curr++
		}
	}
}
// @lc code=end

