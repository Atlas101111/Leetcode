/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode.cn/problems/non-decreasing-subsequences/description/
 *
 * algorithms
 * Medium (52.62%)
 * Likes:    576
 * Dislikes: 0
 * Total Accepted:    114.3K
 * Total Submissions: 217.1K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(input []int, startIndex int)
	dfs = func(input []int, startIndex int) {
		for i := startIndex; i < len(input); i += 1 {
			if len(path) > 0 && path[len(path)-1] > input[i] {
				continue
			}

			flag := false
			for x := startIndex; x < i; x += 1 {
				if input[x] == input[i] {
					flag = true
					break
				}
			}
			if flag {
				continue
			}

			if len(path) >= 1 {
				tempPath := make([]int, len(path), 2*len(path))
				copy(tempPath, path)

				tempPath = append(tempPath, input[i])
				result = append(result, tempPath)
			}

			path = append(path, input[i])
			dfs(input, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return result
}

// @lc code=end

