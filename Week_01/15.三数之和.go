/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	//暴力破解 枚举 a + b = -c  因为要排序,执行超时了 
	/*
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	ret := [][]int{}
	sort.Ints(nums)	//先排个序, 要不然后面没法玩 
	htab := make(map[string]int)
	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := j+1; k < l; k++ {
				tmp := []int{nums[i], nums[j], nums[k]}
				sort.Ints(tmp)
				s := func(n []int) string {
					s := ""
					for _, v := range n {
						s += strconv.Itoa(v)
					}
					return s
				}(tmp)
				_, ok := htab[s]	//去重 
				if  !ok && (nums[k] + nums[j] == -nums[i]) { 
					htab[s]++
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
	*/
	//双指针法
	l := len(nums)
	ret := [][]int{}
	if l == 0 || l < 3{
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < l - 1; i++ {
		if nums[i] > 0 {	//排序之后，第一个数都大于0的话，则数组不存在3数之和为0的情况
			break
		}
		if i > 0 && (nums[i] == nums[i - 1]) {	//去除重复的情况
			continue
		}
		left := i + 1
		right := l - 1
		for left < right {
				if nums[i] + nums[left] + nums[right] == 0 {
					ret = append(ret, []int{nums[i], nums[left], nums[right]})
					for (left < right) && (nums[left] == nums[left + 1]) {
						left++
					}
					for (left < right) && (nums[right] == nums[right -1]) {
						right--
					}
					left++
					right--
				} else if nums[i] + nums[left] + nums[right] > 0 {
					right--
				} else {
					left++
				}
		}
	}
	return ret
}
// @lc code=end

