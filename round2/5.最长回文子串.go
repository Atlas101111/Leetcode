/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.39%)
 * Likes:    6859
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 4.1M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for index, _ := range dp {
		dp[index] = make([]bool, len(s))
	}

	result := s[0:1]
	for length := 1; length <= len(s); length += 1 {
		for i := 0; i <= len(s)-length; i += 1 {
			j := i + length - 1
			if length == 1 {
				dp[i][j] = true
				continue
			}
			if length == 2 {
				if s[j] == s[i] {
					dp[i][j] = true
					if len(result) < length {
						result = s[i : j+1]
					}
				}
				continue
			}

			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if len(result) < length {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}

// @lc code=end

