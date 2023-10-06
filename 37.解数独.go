/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode.cn/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (67.65%)
 * Likes:    1516
 * Dislikes: 0
 * Total Accepted:    186.5K
 * Total Submissions: 275.6K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过填充空格来解决数独问题。
 *
 * 数独的解法需 遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
 *
 *
 * 数独部分空格内已填入了数字，空白格用 '.' 表示。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
 *
 * 输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
 * 解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] 是一位数字或者 '.'
 * 题目数据 保证 输入数独仅有一个解
 *
 *
 *
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte) {

	// used := make([]bool, 10)

	var dfs func(board [][]byte) bool
	dfs = func(board [][]byte) bool {
		for i := 0; i < len(board); i += 1 {
			for j := 0; j < len(board); j += 1 {
				if board[i][j] != '.' {
					continue
				}
				for value := 1; value <= 9; value += 1 {
					if isValid(board, i, j, value) {
						valueString := strconv.Itoa(value)
						board[i][j] = valueString[0]
						if dfs(board) {
							return true
						}
						board[i][j] = "."[0]
					}
				}
			}
			return false
		}
		return true
	}

	dfs(board)
}

func isValid(board [][]byte, i int, j int, value int) bool {

	// 同列
	for x := 1; x <= i; x += 1 {
		boardValue, _ := strconv.Atoi(string(board[i-x][j]))
		if boardValue == value {
			return false
		}
	}
	// 同行
	for x := 1; x <= j; x += 1 {
		boardValue, _ := strconv.Atoi(string(board[i][j-x]))
		if boardValue == value {
			return false
		}
	}

	// 九宫格内
	startLevel := (i / 3) * 3
	startRow := (j / 3) * 3
	for x := startLevel; x < startLevel+3; x += 1 {
		for y := startRow; y < startRow+3; y += 1 {
			boardValue, _ := strconv.Atoi(string(board[x][y]))
			if boardValue == value {
				return false
			}
		}
	}

	return true
}

// @lc code=end

