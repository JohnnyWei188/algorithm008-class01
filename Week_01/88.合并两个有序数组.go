/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	//暴力法
	/*
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
	*/
	//末尾双指针法
	i, j := m - 1, n - 1
	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[k] = nums1[i]
			i--
		}else{
			nums1[k] = nums2[j]
			j--
		}
	}
	
}
// @lc code=end

