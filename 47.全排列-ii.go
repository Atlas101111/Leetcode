/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode.cn/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (65.38%)
 * Likes:    1270
 * Dislikes: 0
 * Total Accepted:    413.2K
 * Total Submissions: 631.9K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var dfs func(input []int, used []bool)
	dfs = func(input []int, used []bool) {
		if len(path) == len(nums) {
			tempPath := make([]int, len(path), 2*len(path))
			copy(tempPath, path)

			result = append(result, tempPath)
			return
		}

		for i := 0; i < len(input); i += 1 {
			if used[i] {
				continue
			}

			if i > 0 && input[i] == input[i-1] && used[i-1] == false {
				continue
			}

			used[i] = true
			path = append(path, input[i])
			dfs(input, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(nums, make([]bool, len(nums), 2*len(nums)))
	return result
}

// @lc code=end

