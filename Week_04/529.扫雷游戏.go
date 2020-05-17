/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	n := len(board)
	m := len(board[0])
	row := click[0]
	col := click[1]
	if board[row][col] == 'M' || board[row][col] == 'X' {
		board[row][col] = 'X'
		return board
	}
	dirs := [][]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	num := 0
	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < n && newCol >= 0 && newCol < m && board[newRow][newCol] == 'M' {
			num++
		}
	}
	if num > 0 {
		board[row][col] =  byte(num + 48)
		return board
	}
	board[row][col] = 'B'
	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < n && newCol >= 0 && newCol < m && board[newRow][newCol] == 'E' {
			nextClick := []int{newRow, newCol}
			updateBoard(board, nextClick)
		}
	}
	return board
}
// @lc code=end

