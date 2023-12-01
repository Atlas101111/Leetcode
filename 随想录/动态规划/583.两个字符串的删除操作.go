/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 *
 * https://leetcode.cn/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (66.22%)
 * Likes:    639
 * Dislikes: 0
 * Total Accepted:    141.4K
 * Total Submissions: 211.4K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
 *
 * 每步 可以删除任意一个字符串中的一个字符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: word1 = "sea", word2 = "eat"
 * 输出: 2
 * 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
 *
 *
 * 示例  2:
 *
 *
 * 输入：word1 = "leetcode", word2 = "etco"
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= word1.length, word2.length <= 500
 * word1 和 word2 只包含小写英文字母
 *
 *
 */

// @lc code=start
func minDistanceUseCommonString(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for index, _ := range dp {
		dp[index] = make([]int, len(word2)+1)

	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	maxCommon := dp[len(word1)][len(word2)]
	result := (len(word1) - maxCommon) + (len(word2) - maxCommon)
	return result
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for index, _ := range dp {
		dp[index] = make([]int, len(word2)+1)
		if index == 0 {
			for i := 0; i <= len(word2); i++ {
				dp[0][i] = i
			}
		}
		dp[index][0] = index
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]+1)
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+2)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// @lc code=end

