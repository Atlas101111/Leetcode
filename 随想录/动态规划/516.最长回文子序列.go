/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 *
 * https://leetcode.cn/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (67.10%)
 * Likes:    1141
 * Dislikes: 0
 * Total Accepted:    208.9K
 * Total Submissions: 311.3K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	result := 1
	dp := make([][]int, len(s)+2)
	for index, _ := range dp {
		dp[index] = make([]int, len(s)+2)
	}

	for i := len(s); i >= 1; i-- {
		for j := i; j <= len(s); j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}

			if s[i-1] == s[j-1] {
				dp[i][j] = 2 + dp[i+1][j-1]
				if dp[i][j] > result {
					result = dp[i][j]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1], dp[i+1][j-1])
			}
		}
	}
	return result
}

// @lc code=end

