/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode.cn/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (75.61%)
 * Likes:    1071
 * Dislikes: 0
 * Total Accepted:    446.3K
 * Total Submissions: 591.1K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: numRows = 5
 * 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 *
 *
 * 示例 2:
 *
 *
 * 输入: numRows = 1
 * 输出: [[1]]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 *
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for index, _ := range dp {
		if index == 0 {
			dp[index] = []int{1}
			continue
		} else if index == 1 {
			dp[index] = []int{1, 1}
			continue
		}
		item := make([]int, index+1)
		for itemIndex, _ := range item {
			if itemIndex == 0 || itemIndex == index {
				item[itemIndex] = 1
			} else {
				item[itemIndex] = dp[index-1][itemIndex] + dp[index-1][itemIndex-1]
			}
		}
		dp[index] = item
	}
	return dp
}

// @lc code=end

