/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	//暴力法
	/*
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		return intersect(nums2, nums1)
	}
	ret := []int{}
	hmap := map[int]int{}
	for _, v := range nums1 {
		hmap[v]++
	}
	for _, v := range nums2 {
		cnt, ok := hmap[v]
		if ok && cnt > 0 {
			ret = append(ret, v)
			hmap[v]--
		}
	}
	return ret
	*/
	//排序+双指针 
	sort.Ints(nums1) //4 4 4 5 9
	sort.Ints(nums2) //4 4 8 9 9
	i, j, k := 0, 0, 0
	for i < len(nums1)  && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			k++
			i++
			j++
		}
	}
	return nums1[0:k]
}

// @lc code=end

