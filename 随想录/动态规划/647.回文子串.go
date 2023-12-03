/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode.cn/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (66.92%)
 * Likes:    1264
 * Dislikes: 0
 * Total Accepted:    307.8K
 * Total Submissions: 459K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
func countSubstrings(s string) int {
	result := 0
	dp := make([][]int, len(s)+1)
	for index, _ := range dp {
		dp[index] = make([]int, len(s)+1)
	}
	for i := 0; i < len(s); i++ {
		dp[i][1] = 1
		result += 1
	}

	for l := 2; l <= len(s); l++ {
		for start := 0; start <= len(s)-l; start++ {
			if s[start] == s[start+l-1] {
				if l == 2 {
					dp[start][l] = 1
					result += 1
				} else if dp[start+1][l-2] == 1 {
					dp[start][l] = 1
					result += 1
				}
			}
		}
	}

	return result
	516
}

// @lc code=end

