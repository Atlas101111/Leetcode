/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.38%)
 * Likes:    410
 * Dislikes: 0
 * Total Accepted:    110.8K
 * Total Submissions: 134.5K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */

// @lc code=start
func totalNQueens(n int) int {
	checkerboard := make([][]bool, n)
	for i := 0; i < n; i += 1 {
		checkerboard[i] = make([]bool, n)
	}

	resultCount := 0

	var dfs func(level, target int) bool
	dfs = func(level, target int) bool {
		if level == target {
			resultCount += 1
			return true
		}

		for i := 0; i < target; i += 1 {
			if isValid(checkerboard, level, i, target) {
				checkerboard[level][i] = true
				// if dfs(level+1, target) {
				// 	return true
				// }
				// 这种写法是用来检测存在性或者整单一解的，只要搜索到树中有一个合法解，就会立刻返回并停止搜索
				dfs(level+1, target)
				checkerboard[level][i] = false
			}
		}

		return false
	}

	dfs(0, n)
	return resultCount
}

func isValid(board [][]bool, i int, j int, n int) bool {
	if i == 0 {
		return true
	}

	// 检查同一列
	for x := 0; x < i; x += 1 {
		if board[x][j] {
			return false
		}
	}

	// 检查左上方
	for x := 1; i-x >= 0 && j-x >= 0; x += 1 {
		if board[i-x][j-x] {
			return false
		}
	}

	for x := 1; i-x >= 0 && j+x <= n-1; x += 1 {
		if board[i-x][j+x] {
			return false
		}
	}

	return true
}

// @lc code=end

