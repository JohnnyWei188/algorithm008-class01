/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	//暴力的说
	/*
	ret := []string{}
	for i := 1; i <= n; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			ret = append(ret, "FizzBuzz")
			continue
		} else if i % 5 == 0 {
			ret = append(ret, "Buzz")
			continue
		} else if  i % 3 == 0 {
			ret = append(ret, "Fizz")
			continue
		}
		ret = append(ret, strconv.Itoa(i))
	}
	return ret */
	//使用散列表保存 对应关系
	hmap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	sortRel := []int{3, 5}	//go 的map 在rangg的时候是无序的, 需要注意
	ret := []string{}
	for i := 1; i <= n; i++ {
		nv := ""
		for _, v := range sortRel {
			if i % v == 0 {
				nv += hmap[v]
			}
		}
		if nv != "" {
			ret = append(ret, nv)
			continue
		}
		ret = append(ret, strconv.Itoa(i))
	}
	return ret
}
// @lc code=end

