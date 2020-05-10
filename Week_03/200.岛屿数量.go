/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	//对着答案敲了一遍的
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {            //遍历岛屿数量
                BFS(grid, i, j)
                result ++
            }
        }
    }
    return result
}

var fx = [4]int{1, 0, -1, 0}     
var fy = [4]int{0, -1, 0, 1}     

func BFS(grid [][]byte, x int, y int){
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0'{ //递归退出临界条件
        return
    }
    grid[x][y] = '0'                  
    for i := 0; i < 4; i++ {         
        BFS(grid,x + fx[i], y + fy[i]) 
    }
}
// @lc code=end

