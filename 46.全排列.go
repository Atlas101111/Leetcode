/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.85%)
 * Likes:    2375
 * Dislikes: 0
 * Total Accepted:    787.2K
 * Total Submissions: 998.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(input []int, used []bool)
	dfs = func(input []int, used []bool) {
		if len(path) == len(input) {
			tempPath := make([]int, len(path), 2*len(path))
			copy(tempPath, path)

			result = append(result, tempPath)
			return
		}

		for i := 0; i < len(input); i += 1 {
			if !used[i] {
				path = append(path, input[i])
				used[i] = true
				dfs(input, used)
				path = path[:len(path)-1]
				used[i] = false
			}

		}
	}

	dfs(nums, make([]bool, len(nums), 2*len(nums)))
	return result
}

// @lc code=end

