/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (80.97%)
 * Likes:    1895
 * Dislikes: 0
 * Total Accepted:    566.3K
 * Total Submissions: 699.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	result = append(result, path)

	var dfs func(input []int, endIndex int)
	dfs = func(input []int, endIndex int) {
		if endIndex >= len(input) {
			return
		}

		currentDigit := input[endIndex]
		preResult := make([][]int, len(result), 2*len(result))

		copy(preResult, result)

		for _, item := range preResult {
			temp := make([]int, len(item), 2*len(item))
			copy(temp, item)
			temp = append(temp, currentDigit)

			result = append(result, temp)
		}

		dfs(input, endIndex+1)
	}

	dfs(nums, 0)
	return result
}

// @lc code=end

