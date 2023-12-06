/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.27%)
 * Likes:    1248
 * Dislikes: 0
 * Total Accepted:    470.4K
 * Total Submissions: 608.7K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=start
var (
	result [][]int
	path   []int
)

func combine(n int, k int) [][]int {
	result, path = make([][]int, 0), make([]int, 0)
	dfs(1, n, k)
	return result
}

func dfs(start, end, targetLen int) {
	if targetLen == len(path) {
		temp := make([]int, targetLen)
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := start; i <= end; i++ {
		if end-i+1 < targetLen-len(path) {
			break
		}

		path = append(path, i)
		dfs(i+1, end, targetLen)
		path = path[:len(path)-1]
	}
}

// @lc code=end

