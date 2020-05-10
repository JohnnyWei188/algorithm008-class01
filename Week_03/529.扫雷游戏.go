/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	rowLen := len(board)	//多少行
	colLen := len(board[0]) //每行的列
	row, col := click[0], click[1] //行 列坐标
	//如果点击的位置是雷的话 直接返回
	if board[row][col] == 'M' || board[row][col] == 'X' {
		board[row][col] = 'X'
		return board
	}
	//定义点击坐标 需要移动的8个方向
	dirs := [][]int{
		{0, -1},	//上
		{0, 1},		//下
		{-1, 0},	//左
		{1, 0},		//右
		{-1, 1},	//左上角
		{1, -1},	//右上角
		{-1, -1},	//左下角
		{1, 1},		//右下角
	}
	num := 0	//记录当前坐标 8个方向 雷的个数
	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newCol >= 0 && 
		newRow < rowLen && newCol < colLen &&
		board[newRow][newCol] == 'M' {
			num++
		} 
	}
	if num > 0 {
		board[row][col] = byte(num + 48)
		return board
	}
	board[row][col] = 'B'
	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newCol >= 0 && 
		newRow < rowLen && newCol < colLen &&
		board[newRow][newCol] == 'E' {
			updateBoard(board, []int{newRow, newCol})
		} 
	}
	return board
 }
// @lc code=end

