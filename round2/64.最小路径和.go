/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.47%)
 * Likes:    1587
 * Dislikes: 0
 * Total Accepted:    522.9K
 * Total Submissions: 750K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for index, _ := range dp {
		dp[index] = make([]int, len(grid[0]))
	}

	for line := 0; line < len(grid); line++ {
		for col := 0; col < len(grid[0]); col++ {
			if line == 0 && col == 0 {
				dp[line][col] = grid[line][col]
				continue
			}
			if line == 0 {
				dp[line][col] = dp[line][col-1] + grid[line][col]
				continue
			}
			if col == 0 {
				dp[line][col] = dp[line-1][col] + grid[line][col]
				continue
			}
			dp[line][col] = min(dp[line-1][col]+grid[line][col], dp[line][col-1]+grid[line][col])
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

