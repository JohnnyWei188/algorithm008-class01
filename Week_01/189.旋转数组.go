/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	/* 暴力法
	l := len(nums)
	if l == 0 || k <= 0 || k > l {
		return
	}
	*/
	/* 
	for i := 0; i < k; i++ {
		mvD := nums[l-1]   //需要挪动的数
		for j := l - 1; j > 0; j-- {  //所有数据往后移动一个位置
			nums[j] = nums[j-1]
		}
		nums[0] = mvD   //放到第一位
	}
	*/ //3次旋转法 [1, 2, 3, 4, 5, 6, 7]
	l := len(nums)
	k %= l   //注意这里, 需要移动的个数其实就是 k % len(nums), 比判定k和len的长度好点
	reverse(nums, 0, l - 1)  //[7,6,5,4,3,2,1]
	reverse(nums, 0, k - 1)  //[5,6,7,4,3,2,1]
	reverse(nums, k, l - 1)  //[7,6,5,1,2,3,4]
	
}
func reverse(nums []int, start, end int){
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
// @lc code=end

